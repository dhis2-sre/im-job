package job

import (
	"github.com/dhis2-sre/im-job/pkg/model"
)

type Repository interface {
	List() ([]*model.Job, error)
	FindById(id uint) (*model.Job, error)
}

func ProvideRepository() Repository {
	return &repository{}
}

type repository struct {
}

func (r repository) List() ([]*model.Job, error) {
	return []*model.Job{
		{
			ID:          1,
			Name:        "env",
			Description: "Show environment variables",
			JobType:     "database",
			Script:      "/scripts/database/env.sh",
		},
		{
			ID:          2,
			Name:        "psql",
			Description: "psql show tables",
			JobType:     "database",
			Script:      "/scripts/database/psql.sh",
		},
		{
			ID:          3,
			Name:        "save",
			Description: "Save database",
			JobType:     "database",
			Script:      "/scripts/database/save.sh",
		},
		{
			ID:          4,
			Name:        "saveas",
			Description: "Save copy of database",
			JobType:     "database",
			Script:      "/scripts/database/saveAs.sh",
		},
	}, nil
}

func (r repository) FindById(id uint) (*model.Job, error) {
	list, _ := r.List()

	job := filter(list, func(job *model.Job) bool {
		return job.ID == id
	})
	if job != nil {
		return job, nil
	}

	return nil, nil
}

func filter(jobs []*model.Job, test func(job *model.Job) bool) *model.Job {
	for _, job := range jobs {
		if test(job) {
			return job
		}
	}
	return nil
}
