package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/conf"
)

type Data struct {
}

func NewData(logger log.Logger, c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
