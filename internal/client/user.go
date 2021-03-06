package client

import (
	"github.com/dhis2-sre/im-job/pkg/config"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
)

func ProvideUser(config config.Config) userClient.Client {
	return userClient.ProvideClient(config.UserService.Host, config.UserService.BasePath)
}
