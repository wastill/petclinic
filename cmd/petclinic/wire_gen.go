// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wastill/petclinic/internal/conf"
	"github.com/wastill/petclinic/internal/server"
	"github.com/wastill/petclinic/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	petClinicRpcService := service.NewPetClinicRPCService()
	grpcServer := server.NewGRPCServer(confServer, petClinicRpcService, logger)
	petClinicRestService := service.NewPetClinicRestService()
	httpServer := server.NewHTTPServer(confServer, petClinicRestService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
