package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/customer/job/v1"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/app/customer/job/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(_ log.Logger, c *conf.Server, customer *service.CustomerService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	apiV1.RegisterCustomerServer(srv, customer)
	return srv
}
