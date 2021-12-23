package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Customer struct {
	Name string
}

type CustomerRepo interface {
	Consume(ctx context.Context, name string) (string, error)
}

type CustomerUseCase struct {
	log  *log.Helper
	repo CustomerRepo
}

func NewCustomerUseCase(logger log.Logger, repo CustomerRepo) *CustomerUseCase {
	return &CustomerUseCase{
		log:  log.NewHelper(log.With(logger, "customer", "biz/customer")),
		repo: repo,
	}
}

func (c *CustomerUseCase) Consume(ctx context.Context, name string) string {
	result, err := c.repo.Consume(ctx, name)
	if err != nil {
		c.log.Error(err)
	}
	return result
}
