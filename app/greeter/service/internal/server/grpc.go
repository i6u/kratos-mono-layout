package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	apiV1 "github.com/i6u/kratos-mono-layout/api/greeter/service/v1"
	"github.com/i6u/kratos-mono-layout/app/greeter/service/internal/conf"
	"github.com/i6u/kratos-mono-layout/app/greeter/service/internal/service"
	"github.com/i6u/kratos-mono-layout/pkg/util/srv"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(logger log.Logger, c *conf.Server, greeter *service.GreeterService) *grpc.Server {
	return srv.NewGRPCServer(
		logger,
		c.Grpc.Network,
		c.Grpc.Address,
		c.Grpc.Timeout.AsDuration(),
		func(srv *grpc.Server) {
			apiV1.RegisterGreeterServer(srv, greeter)
		},
	)
}
