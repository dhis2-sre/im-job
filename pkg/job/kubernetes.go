package job

import (
	"context"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"go.mozilla.org/sops/v3/cmd/sops/formats"
	"go.mozilla.org/sops/v3/decrypt"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesService interface {
	RunJob(name, script, namespace string, payload map[string]string, configuration *models.ClusterConfiguration) (string, error)
}

func ProvideKubernetesService() KubernetesService {
	return &kubernetesService{}
}

type kubernetesService struct{}

func (k kubernetesService) RunJob(name, script, namespace string, payload map[string]string, configuration *models.ClusterConfiguration) (string, error) {
	image := "dhis2/im-job-runner"

	client, err := k.getClient(configuration)
	if err != nil {
		return "", err
	}

	jobs := client.BatchV1().Jobs(namespace)
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    name,
							Image:   image,
							Command: []string{"/scripts/hello.sh"},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err = jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		return "", err
	}

	return name, nil
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
