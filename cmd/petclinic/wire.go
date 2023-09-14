//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/wastill/petclinic/internal/biz"
	"github.com/wastill/petclinic/internal/conf"
	"github.com/wastill/petclinic/internal/data"
	"github.com/wastill/petclinic/internal/server"
	"github.com/wastill/petclinic/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	//panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
