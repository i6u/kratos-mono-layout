package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/i6u/kratos-mono-layout/app/greeter/service/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

type Data struct {
	// TODO database client
}

func NewData(logger log.Logger, c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
