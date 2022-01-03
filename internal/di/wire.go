//+build wireinject

package di

import (
	"github.com/dhis2-sre/im-job/internal/handler"
	"github.com/dhis2-sre/im-job/pkg/config"
	"github.com/google/wire"
)

type Environment struct {
	Config                   config.Config
	AuthenticationMiddleware handler.AuthenticationMiddleware
}

func ProvideEnvironment(
	config config.Config,
	authenticationMiddleware handler.AuthenticationMiddleware,
) Environment {
	return Environment{
		config,
		authenticationMiddleware,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		config.ProvideConfig,

		//		client.ProvideUser,

		handler.ProvideAuthentication,

		ProvideEnvironment,
	)
	return Environment{}
}
