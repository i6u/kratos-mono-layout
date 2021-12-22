package service

import (
	"context"

	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/greeter/service/v1"
)

func (g *GreeterService) SayHi(ctx context.Context, req *apiV1.SayHiRequest) (*apiV1.SayHiReply, error) {
	msg := g.guc.SayHi(ctx, req.Name)
	return &apiV1.SayHiReply{
		Message: msg,
	}, nil
}
