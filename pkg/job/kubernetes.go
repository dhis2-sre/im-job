package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/lithammer/shortuuid/v3"
	"go.mozilla.org/sops/v3/cmd/sops/formats"
	"go.mozilla.org/sops/v3/decrypt"
	batchv1 "k8s.io/api/batch/v1"
	v1core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
)

type KubernetesService interface {
	RunJob(name, script, namespace string, payload map[string]string, configuration *models.ClusterConfiguration) (string, error)
	JobStatus(rid string, namespace string, configuration *models.ClusterConfiguration) (*batchv1.JobStatus, error)
	Executor(configuration *models.ClusterConfiguration, fn func(client *kubernetes.Clientset) error) (error, error)
}

func ProvideKubernetesService() KubernetesService {
	return &kubernetesService{}
}

type kubernetesService struct{}

func (k kubernetesService) RunJob(name, script, namespace string, payload map[string]string, configuration *models.ClusterConfiguration) (string, error) {
	image := "dhis2/im-job-runner"
	uuid := shortuuid.New()
	runId := name + "-" + strings.ToLower(uuid)

	client, err := k.getClient(configuration)
	if err != nil {
		return "", err
	}

	jobs := client.BatchV1().Jobs(namespace)
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      runId,
			Namespace: namespace,
			Labels: map[string]string{
				"runId": runId,
			},
		},
		Spec: batchv1.JobSpec{
			Template: v1core.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"runId": runId,
					},
				},
				Spec: v1core.PodSpec{
					Containers: []v1core.Container{
						{
							Name:    runId,
							Image:   image,
							Command: []string{"/scripts/hello.sh"},
						},
					},
					RestartPolicy: v1core.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err = jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		return "", err
	}

	return runId, nil
}

func (k kubernetesService) JobStatus(runId string, namespace string, configuration *models.ClusterConfiguration) (*batchv1.JobStatus, error) {
	client, err := k.getClient(configuration)
	if err != nil {
		return &batchv1.JobStatus{}, err
	}

	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("runId=%s", runId),
	}

	list, err := client.BatchV1().Jobs(namespace).List(context.TODO(), listOptions)
	if err != nil {
		log.Printf("%+v", err)
		return &batchv1.JobStatus{}, err
	}

	jobs := list.Items

	if len(jobs) < 1 {
		return &batchv1.JobStatus{}, errors.New("no jobs found with runId: " + runId)
	}

	if len(jobs) > 1 {
		return &batchv1.JobStatus{}, errors.New("multiple jobs found with runId: " + runId)
	}

	return &jobs[0].Status, nil
}

func (k kubernetesService) Executor(configuration *models.ClusterConfiguration, fn func(client *kubernetes.Clientset) error) (error, error) {
	client, err := k.getClient(configuration)
	if err != nil {
		return err, nil
	}
	return nil, fn(client)
}

func (k kubernetesService) getClient(configuration *models.ClusterConfiguration) (*kubernetes.Clientset, error) {
	var restClientConfig *rest.Config
	if len(configuration.KubernetesConfiguration) > 0 {
		configurationInCleartext, err := k.decrypt(configuration.KubernetesConfiguration, "yaml")
		if err != nil {
			return nil, err
		}

		config, err := clientcmd.NewClientConfigFromBytes(configurationInCleartext)
		if err != nil {
			return nil, err
		}

		restClientConfig, err = config.ClientConfig()
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		restClientConfig, err = clientcmd.BuildConfigFromFlags("", "")
		if err != nil {
			return nil, err
		}
	}

	client, err := kubernetes.NewForConfig(restClientConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (k kubernetesService) decrypt(data []byte, format string) ([]byte, error) {
	kubernetesConfigurationCleartext, err := decrypt.DataWithFormat(data, formats.FormatFromString(format))
	if err != nil {
		return nil, err
	}
	return kubernetesConfigurationCleartext, nil
}
