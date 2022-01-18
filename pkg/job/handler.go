package job

import (
	"bufio"
	"github.com/dhis2-sre/im-job/internal/apperror"
	"github.com/dhis2-sre/im-job/internal/handler"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func ProvideHandler(userClient userClient.Client, jobService Service) Handler {
	return Handler{
		userClient,
		jobService,
	}
}

type Handler struct {
	userClient userClient.Client
	jobService Service
}

// List jobs
// swagger:route GET /jobs listJobs
//
// List jobs
//
// Security:
//   oauth2:
//
// responses:
//   200: []Job
//   401: Error
//   403: Error
//   415: Error
func (h Handler) List(c *gin.Context) {
	jobs, err := h.jobService.List()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, jobs)
}

// FindById job
// swagger:route GET /jobs/{id} findJob
//
// Find job by id
//
// Security:
//   oauth2:
//
// responses:
//   200: Job
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	job, err := h.jobService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, job)
}

type RunJobRequest struct {
	GroupID  uint              `json:"groupId" binding:"required"`
	TargetID uint              `json:"targetId" binding:"required"`
	Payload  map[string]string `json:"payload"`
}

type RunJobResponse struct {
	RunId string `json:"runId"`
}

// Run job
// swagger:route POST /jobs/{id}/run runJob
//
// Run job
//
// Security:
//   oauth2:
//
// responses:
//   200: RunJobResponse
//   401: Error
//   400: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var request RunJobRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	job, err := h.jobService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.userClient.FindGroupById(token, request.GroupID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	runId, err := h.jobService.Run(uint(id), job.JobType, request.TargetID, group, request.Payload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, RunJobResponse{runId})
}

type StatusRequest struct {
	GroupID uint `json:"groupId" binding:"required"`
}

// Status job
// swagger:route GET /jobs/running/{runId}/status jobStatus
//
// Job status
//
// Security:
//   oauth2:
//
// responses:
//   200: Job
//   400: Error
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Status(c *gin.Context) {
	runId := c.Param("runId")
	if runId == "" {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var request StatusRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.userClient.FindGroupById(token, request.GroupID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	status, err := h.jobService.Status(runId, group)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, status)
}

type LogsRequest struct {
	GroupID uint `json:"groupId" binding:"required"`
}

// Logs job
// swagger:route GET /jobs/running/{runId}/logs jobLogs
//
// Job logs
//
// Security:
//   oauth2:
//
// responses:
//   200:
//   400: Error
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Logs(c *gin.Context) {
	runId := c.Param("runId")
	if runId == "" {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var request StatusRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.userClient.FindGroupById(token, request.GroupID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	readCloser, err := h.jobService.Logs(runId, group)
	if err != nil {
		_ = c.Error(err)
		return
	}

	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			_ = c.Error(err)
		}
	}(readCloser)

	bufferedReader := bufio.NewReader(readCloser)

	c.Stream(func(writer io.Writer) bool {
		readBytes, err := bufferedReader.ReadBytes('\n')
		if err != nil {
			return false
		}

		_, err = writer.Write(readBytes)
		return err == nil
	})

	c.Status(http.StatusOK)
}
