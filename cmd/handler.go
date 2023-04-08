package cmd

import (
	"go-skeleton/internal/handler"
	db "go-skeleton/internal/infrastructure/db/gorm"
	repository "go-skeleton/internal/infrastructure/repository/example"
	usecase "go-skeleton/internal/usecase/example"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func (s server) RegisterHandler() {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	// init database conn
	db := db.NewDB()

	// init redis conn
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// init repo
	exampleRepo := repository.NewExampleRepo(db)

	// init usecase
	exampleUsecase := usecase.NewExampleUsecase(exampleRepo, rdb)

	// init handler
	exampleGroup := s.router.Group("/example")
	handler.RegisterExampleHandler(exampleGroup, exampleUsecase)
}
