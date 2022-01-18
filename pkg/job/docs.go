package job

import (
	"github.com/dhis2-sre/im-job/pkg/model"
)

// swagger:response Error
type _ struct {
	//in: body
	Body string
}

// swagger:response Job
type _ struct {
	//in: body
	Body model.Job
}

// swagger:response RunJobResponse
type _ struct {
	//in: body
	Body RunJobResponse
}

//swagger:parameters findJob runJob
type _ struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}

//swagger:parameters jobStatus jobLogs
type _ struct {
	// in: path
	// required: true
	ID uint `json:"runId"`
}

// swagger:parameters runJob
type _ struct {
	// Run job request body parameter
	// in: body
	// required: true
	Body RunJobRequest
}

// swagger:parameters jobLogs
type _ struct {
	// Logs request body parameter
	// in: body
	// required: true
	Body LogsRequest
}
