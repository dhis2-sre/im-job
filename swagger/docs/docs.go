package docs

import (
	"github.com/dhis2-sre/im-job/pkg/job"
	"github.com/dhis2-sre/im-job/pkg/model"
)

// swagger:response
type Error struct {
	// The error message
	//in: body
	Message string
}

// swagger:response
type Job struct {
	//in: body
	Job model.Job
}

//swagger:parameters findJob runJob
type IdParam struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}

// swagger:parameters runJob
type _ struct {
	// Run job request body parameter
	// in: body
	// required: true
	Body job.RunJobRequest
}