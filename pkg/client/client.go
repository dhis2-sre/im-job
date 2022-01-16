package client

import (
	"context"
	"github.com/dhis2-sre/im-job/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-job/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	Run(token string, id uint, body *models.RunJobRequest) (string, error)
}

func ProvideClient(host string, basePath string) Client {
	transport := httptransport.New(host, basePath, nil)
	service := operations.New(transport, strfmt.Default)
	return &cli{service}
}

type cli struct {
	clientService operations.ClientService
}

func (c cli) Run(token string, id uint, body *models.RunJobRequest) (string, error) {
	params := &operations.RunJobParams{ID: uint64(id), Body: body, Context: context.Background()}
	clientAuthInfoWriter := httptransport.BearerToken(token)

	response, err := c.clientService.RunJob(params, clientAuthInfoWriter)
	if err != nil {
		return "", err
	}

	return response.GetPayload().RunID, nil
}
