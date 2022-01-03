//go:build wireinject
// +build wireinject

package di

import (
	"github.com/dhis2-sre/im-job/internal/client"
	"github.com/dhis2-sre/im-job/internal/handler"
	"github.com/dhis2-sre/im-job/pkg/config"
	"github.com/dhis2-sre/im-job/pkg/job"
	"github.com/google/wire"
)

type Environment struct {
	Config                   config.Config
	AuthenticationMiddleware handler.AuthenticationMiddleware
	JobHandler               job.Handler
}

func ProvideEnvironment(
	config config.Config,
	authenticationMiddleware handler.AuthenticationMiddleware,
	jobHandler job.Handler,
) Environment {
	return Environment{
		config,
		authenticationMiddleware,
		jobHandler,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		config.ProvideConfig,

		client.ProvideUser,

		handler.ProvideAuthentication,

		job.ProvideKubernetesService,
		job.ProvideRepository,
		job.ProvideService,
		job.ProvideHandler,

		ProvideEnvironment,
	)
	return Environment{}
}
