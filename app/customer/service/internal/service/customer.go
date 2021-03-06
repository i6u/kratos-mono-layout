package service

import (
	"context"

	apiV1 "github.com/i6u/kratos-mono-layout/api/customer/service/v1"
)

func (c *CustomerService) Consume(ctx context.Context, req *apiV1.ConsumeRequest) (*apiV1.ConsumeReply, error) {
	msg := c.cuc.Consume(ctx, req.Name)
	return &apiV1.ConsumeReply{
		Message: msg,
	}, nil
}
