package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	apiV1 "github.com/i6u/kratos-mono-layout/api/customer/service/v1"
	"github.com/i6u/kratos-mono-layout/app/customer/service/internal/conf"
	"github.com/i6u/kratos-mono-layout/app/customer/service/internal/service"
	"github.com/i6u/kratos-mono-layout/pkg/util/srv"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(logger log.Logger, c *conf.Server, customer *service.CustomerService) *http.Server {
	return srv.NewHTTPServer(
		logger,
		c.Http.Network,
		c.Http.Address,
		c.Grpc.Timeout.AsDuration(),
		func(srv *http.Server) {
			apiV1.RegisterCustomerHTTPServer(srv, customer)
		},
	)
}
