package usecase

import (
	"context"
	"go-skeleton/internal/entity"
	repository "go-skeleton/internal/infrastructure/repository/example"
	"go-skeleton/internal/request"
)

func NewExampleUsecase(exampleRepo repository.ExampleRepository) ExampleUsecase {
	return exampleUsecase{
		repo: exampleRepo,
	}
}

type ExampleUsecase interface {
	CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error)
}

type exampleUsecase struct {
	repo repository.ExampleRepository
}

func (s exampleUsecase) CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error) {
	example := entity.Example{
		Name: req.Name,
	}
	return s.repo.CreateExample(ctx, &example)
}
