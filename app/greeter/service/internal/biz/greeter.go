package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Name string
}

type GreeterRepo interface {
	GetGreeter(ctx context.Context, name string) (*Greeter, error)
}

type GreeterUseCase struct {
	log  *log.Helper
	repo GreeterRepo
}

func NewGreeterUseCase(logger log.Logger, repo GreeterRepo) *GreeterUseCase {
	return &GreeterUseCase{
		log:  log.NewHelper(log.With(logger, "greeter", "biz/greeter")),
		repo: repo,
	}
}

func (g *GreeterUseCase) SayHi(ctx context.Context, name string) string {
	greeter, err := g.repo.GetGreeter(ctx, name)
	if err != nil {
		return ""
	}
	return "Gutter " + greeter.Name + ": 'Say hi!'"
}
