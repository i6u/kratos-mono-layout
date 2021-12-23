package reg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	etcdClient "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func initEtcdClient(endpoints []string, timeout time.Duration) (*etcdClient.Client, func()) {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	})
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		cli.Close()
	}
	return cli, cleanup
}

func NewRegistrar(endpoints []string, timeout time.Duration) (registry.Registrar, func()) {
	cli, cleanup := initEtcdClient(endpoints, timeout)
	return etcd.New(cli), cleanup
}

func NewDiscovery(endpoints []string, timeout time.Duration) (registry.Discovery, func()) {
	cli, cleanup := initEtcdClient(endpoints, timeout)
	return etcd.New(cli), cleanup
}

func NewServiceClientConn(_ log.Logger, r registry.Discovery, name string) (*grpc.ClientConn, error) {
	conn, err := _grpc.DialInsecure(
		context.Background(),
		_grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", name)),
		_grpc.WithTimeout(75*time.Second),
		_grpc.WithDiscovery(r),
		_grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			//logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
