package repository

import (
	"context"
	"go-skeleton/internal/entity"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewExampleRepo(db *gorm.DB) ExampleRepository {
	repo := &exampleRepository{
		db:          db,
		autoMigrate: viper.GetBool("postgres.auto-migrate"),
	}

	if repo.autoMigrate {
		repo.db.AutoMigrate(&entity.Example{})
	}

	return repo
}

type ExampleRepository interface {
	CreateExample(ctx context.Context, entity *entity.Example) (*entity.Example, error)
}

type exampleRepository struct {
	db          *gorm.DB
	autoMigrate bool
}

func (repo *exampleRepository) CreateExample(ctx context.Context, enity *entity.Example) (*entity.Example, error) {
	return nil, nil
}
