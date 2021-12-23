package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/biz"
)

type customerRepo struct {
	log *log.Helper
}

func NewCustomerRepo(logger log.Logger) biz.CustomerRepo {
	return &customerRepo{
		log: log.NewHelper(log.With(logger, "customer", "data/customer")),
	}
}

func (c *customerRepo) Consume(ctx context.Context, names []string) (string, error) {
	return "", nil
}
