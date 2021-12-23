// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/biz"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/data"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/server"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/service"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
