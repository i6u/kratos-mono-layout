package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Name string
}

type GreeterRepo interface {
	GetGreeter(ctx context.Context) (*Greeter, error)
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
	greeter, err := g.repo.GetGreeter(ctx)
	if err != nil {
		return "not found"
	}
	return fmt.Sprintf("%s say : 'Hi %s!'", greeter.Name, name)
}
