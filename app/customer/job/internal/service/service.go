package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/customer/job/v1"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/biz"
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
