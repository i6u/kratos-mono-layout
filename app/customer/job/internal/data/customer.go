package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/biz"

	greeterV1 "github.com/i6u/kratos-multiple-module-layout/api/greeter/service/v1"
)

type customerRepo struct {
	log  *log.Helper
	data *Data
}

func NewCustomerRepo(logger log.Logger, data *Data) biz.CustomerRepo {
	return &customerRepo{
		log:  log.NewHelper(log.With(logger, "customer", "data/customer")),
		data: data,
	}
}

func (c *customerRepo) Consume(ctx context.Context, name string) (string, error) {
	reply, err := c.data.greeter.SayHi(ctx, &greeterV1.SayHiRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return reply.Message, nil
}
