package job

import (
	"github.com/dhis2-sre/im-job/pkg/config"
	"github.com/dhis2-sre/im-job/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

type Service interface {
	List() ([]*model.Job, error)
	FindById(id uint) (*model.Job, error)
	Run(id uint, group *models.Group, payload map[string]string) (string, error)
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
	job, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}

	rid, err := s.kubernetes.RunJob(job.Name, job.Script, group.Name, payload, group.ClusterConfiguration)
	if err != nil {
		return "", err
	}

	return rid, nil
}
