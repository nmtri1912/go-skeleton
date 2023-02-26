package cmd

import (
	"context"
	"go-skeleton/internal/handler"
	repo "go-skeleton/internal/repo/example"
	service "go-skeleton/internal/service/example"
	"go-skeleton/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer() Server {
	return server{
		router: gin.New(),
	}
}

type Server interface {
	Start()
}

type server struct {
	router *gin.Engine
}

func (s server) Start() {
	// source config to viper
	utils.LoadConfiguration()

	// register handler
	s.RegisterHandler()

	// start server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	default:
	}

	log.Println("Server exiting")
}

func (s server) RegisterHandler() {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	// init database conn
	db := utils.NewDB()

	// init repo
	exampleRepo := repo.NewExampleRepo(db)

	// init service
	exampleService := service.NewExampleService(exampleRepo)

	// init handler
	exampleGroup := s.router.Group("/example")
	handler.RegisterExampleHandler(exampleGroup, exampleService)
}
