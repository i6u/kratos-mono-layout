package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	apiV1 "github.com/i6u/kratos-mono-layout/api/customer/service/v1"
	"github.com/i6u/kratos-mono-layout/app/customer/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewCustomerService)

type CustomerService struct {
	apiV1.UnimplementedCustomerServer

	log *log.Helper
	cuc *biz.CustomerUseCase
}

func NewCustomerService(logger log.Logger, cuc *biz.CustomerUseCase) *CustomerService {
	return &CustomerService{
		log: log.NewHelper(log.With(logger, "customer", "service/customer")),
		cuc: cuc,
	}
}
