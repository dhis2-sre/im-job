package job

import (
	"github.com/dhis2-sre/im-job/internal/apperror"
	"github.com/dhis2-sre/im-job/internal/handler"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
	"github.com/gin-gonic/gin"
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
	GroupID uint              `json:"groupId" binding:"required"`
	Payload map[string]string `json:"payload" binding:"required"`
}

// Run job
// swagger:route GET /jobs/{id}/run runJob
//
// Run job
//
// Security:
//   oauth2:
//
// responses:
//   200: Job
//   401: Error
//   403: Error
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

	group, err := h.userClient.FindGroupById(token, request.GroupID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rid, err := h.jobService.Run(uint(id), group, request.Payload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rid)
}
