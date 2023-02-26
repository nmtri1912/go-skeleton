package service

import (
	"context"
	"go-skeleton/internal/entity"
	repo "go-skeleton/internal/repo/example"
	"go-skeleton/internal/request"
)

func NewExampleService(exampleRepo repo.ExampleRepository) ExampleService {
	return exampleService{
		repo: exampleRepo,
	}
}

type ExampleService interface {
	CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error)
}

type exampleService struct {
	repo repo.ExampleRepository
}

func (s exampleService) CreateExample(ctx context.Context, req *request.ExampleRequest) (*entity.Example, error) {
	example := entity.Example{
		Name: req.Name,
	}
	return s.repo.CreateExample(ctx, &example)
}
