package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/dhis2-sre/im-job/pkg/config"
	"github.com/dhis2-sre/im-job/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"io"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Service interface {
	List() ([]*model.Job, error)
	FindById(id uint) (*model.Job, error)
	Run(id uint, group *models.Group, payload map[string]string) (string, error)
	Status(rid string, group *models.Group) (batchv1.JobStatus, error)
	Logs(runId string, group *models.Group) (io.ReadCloser, error)
}

func ProvideService(c config.Config, repository Repository, kubernetes KubernetesService) Service {
	return &service{c, repository, kubernetes}
}

type service struct {
	c          config.Config
	repository Repository
	kubernetes KubernetesService
}

func (s service) List() ([]*model.Job, error) {
	list, err := s.repository.List()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s service) FindById(id uint) (*model.Job, error) {
	job, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s service) Run(id uint, group *models.Group, payload map[string]string) (string, error) {
	// TODO: Merge payload with hostname of dhis2 and postgresql
	// And credentials for postgresql... And perhaps admin/district... And more?

	job, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}

	name := fmt.Sprintf("%s-%s", job.Name, job.JobType)
	runId, err := s.kubernetes.RunJob(name, job.Script, group.Name, payload, group.ClusterConfiguration)
	if err != nil {
		return "", err
	}

	return runId, nil
}

// TODO: Don't return batchv1.JobStatus, the handler shouldn't know about batchv1.JobStatus
func (s service) Status(runId string, group *models.Group) (batchv1.JobStatus, error) {
	status, err := s.kubernetes.JobStatus(runId, group.Name, group.ClusterConfiguration)
	if err != nil {
		return batchv1.JobStatus{}, err
	}

	return status, nil
}

func (s service) Logs(runId string, group *models.Group) (io.ReadCloser, error) {
	var read io.ReadCloser

	err, fnErr := s.kubernetes.Executor(group.ClusterConfiguration, func(client *kubernetes.Clientset) error {
		pod, err := s.getPod(client, runId)
		if err != nil {
			return err
		}

		podLogOptions := v1.PodLogOptions{
			Follow: true,
		}

		readCloser, err := client.
			CoreV1().
			Pods(pod.Namespace).
			GetLogs(pod.Name, &podLogOptions).
			Stream(context.TODO())
		read = readCloser
		return err
	})

	if err != nil {
		return nil, err
	}

	if fnErr != nil {
		return nil, fnErr
	}

	return read, nil
}

// TODO: This method should be moved down to KubernetesService
func (s service) getPod(client *kubernetes.Clientset, runId string) (v1.Pod, error) {
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("runId=%s", runId),
	}

	podList, err := client.CoreV1().Pods("").List(context.TODO(), listOptions)
	if err != nil {
		return v1.Pod{}, err
	}

	if len(podList.Items) > 1 {
		return v1.Pod{}, errors.New("multiple pods found")
	}

	return podList.Items[0], nil
}
