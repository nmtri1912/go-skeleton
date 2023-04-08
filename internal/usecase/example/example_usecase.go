package usecase

import (
	"context"
	"go-skeleton/internal/entity"
	repository "go-skeleton/internal/infrastructure/repository/example"
	"go-skeleton/internal/request"

	"github.com/go-redis/redis/v8"
)

func NewExampleUsecase(exampleRepo repository.ExampleRepository, rdb *redis.Client) ExampleUsecase {
	return exampleUsecase{
		repo:  exampleRepo,
		redis: rdb,
	}
}

type ExampleUsecase interface {
	CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error)
}

type exampleUsecase struct {
	repo  repository.ExampleRepository
	redis *redis.Client
}

func (s exampleUsecase) CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error) {
	example := entity.Example{
		Name: req.Name,
	}
	return s.repo.CreateExample(ctx, &example)
}
