package service

import (
	"context"

	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/customer/job/v1"
)

func (c *CustomerService) Consume(ctx context.Context, req *apiV1.ConsumeRequest) (*apiV1.ConsumeReply, error) {
	names := req.Name
	msg := c.cuc.Consume(ctx, names)
	return &apiV1.ConsumeReply{
		Message: msg,
	}, nil
}
