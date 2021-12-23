package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/pkg/util/reg"

	greeterV1 "github.com/i6u/kratos-multiple-module-layout/api/greeter/service/v1"
)

var ProviderSet = wire.NewSet(NewData, NewCustomerRepo, NewDiscovery, NewGreeterClient)

type Data struct {
	greeter greeterV1.GreeterClient
}

func NewData(logger log.Logger, c *conf.Data, greeter greeterV1.GreeterClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		greeter: greeter,
	}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) (registry.Discovery, func()) {
	return reg.NewDiscovery(conf.Endpoint, conf.Timeout.AsDuration())
}

func NewGreeterClient(logger log.Logger, r registry.Discovery) greeterV1.GreeterClient {
	conn, err := reg.NewServiceClientConn(logger, r, "greeter.service.grpc")
	if err != nil {
		panic(err)
	}
	return greeterV1.NewGreeterClient(conn)
}
