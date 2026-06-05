package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mrdyuke/shorty/config"
	"github.com/mrdyuke/shorty/internal/controller"
	"github.com/mrdyuke/shorty/internal/pool"
	"github.com/mrdyuke/shorty/internal/repo"
	"github.com/mrdyuke/shorty/internal/server"
	"github.com/mrdyuke/shorty/internal/usecase"
)

func main() {
	// == config ==
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("config failed", "error", err)
		return
	}
	slog.Info("app config loaded")

	// == router ==
	if cfg.IsRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// == connection pool ==
	pool, err := pool.NewPostgresPool()
	if err != nil {
		slog.Error("pool failed", "error", err)
		return
	}
	defer pool.Close()
	slog.Info("connection pool loaded")

	// == dependencies ==
	repo := repo.NewURLPostgresRepo(pool)
	usecase := usecase.NewURLUseCase(repo)
	routes := controller.NewURLController(usecase)
	slog.Info("dependencies loaded")

	// == routes ==
	router.GET("/health", controller.HealthCheck)
	router.POST("/short", routes.ShortenURL)

	// == http server ==
	if err := server.RunServer(cfg.ServerPort, router); err != nil {
		slog.Error("server failed", "error", err)
	}
}
