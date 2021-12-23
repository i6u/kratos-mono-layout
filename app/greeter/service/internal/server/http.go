package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	apiV1 "github.com/i6u/kratos-multiple-module-layout/api/greeter/service/v1"
	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/conf"
	"github.com/i6u/kratos-multiple-module-layout/app/greeter/service/internal/service"
	"github.com/i6u/kratos-multiple-module-layout/pkg/util/srv"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(logger log.Logger, c *conf.Server, greeter *service.GreeterService) *http.Server {
	return srv.NewHTTPServer(
		logger,
		c.Http.Network,
		c.Http.Addr,
		c.Grpc.Timeout.AsDuration(),
		func(srv *http.Server) {
			apiV1.RegisterGreeterHTTPServer(srv, greeter)
		},
	)
}
