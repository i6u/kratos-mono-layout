// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/service/internal/biz"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/service/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/service/internal/data"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/service/internal/server"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/service/internal/service"
)

// Injectors from wire.go:

func initApp(registry *conf.Registry, confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery, cleanup := data.NewDiscovery(registry)
	greeterClient := data.NewGreeterClient(logger, discovery)
	dataData, cleanup2, err := data.NewData(logger, confData, greeterClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	customerRepo := data.NewCustomerRepo(logger, dataData)
	customerUseCase := biz.NewCustomerUseCase(logger, customerRepo)
	customerService := service.NewCustomerService(logger, customerUseCase)
	httpServer := server.NewHTTPServer(logger, confServer, customerService)
	grpcServer := server.NewGRPCServer(logger, confServer, customerService)
	registrar, cleanup3 := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}