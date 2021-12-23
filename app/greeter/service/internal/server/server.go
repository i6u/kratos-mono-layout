package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/pkg/util/reg"
)

var ProviderSet = wire.NewSet(NewRegistrar, NewHTTPServer, NewGRPCServer)

func NewRegistrar(conf *conf.Registry) (registry.Registrar, func()) {
	return reg.NewRegistrar(conf.Endpoint, conf.Timeout.AsDuration())
}
