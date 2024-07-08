package app

import (
	"fmt"
	"os"

	"log/slog"

	"github.com/labstack/echo/v4"

	"url-shortener/internal/config"
	v1 "url-shortener/internal/controller/http/v1"
	"url-shortener/internal/repo"
	"url-shortener/internal/service"
	"url-shortener/pkg/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	// ssoClient, err := ssogrpc.New(
	// 	context.Background(),
	// 	log,
	// 	cfg.Clients.SSO.Address,
	// 	cfg.Clients.SSO.Timeout,
	// 	cfg.Clients.SSO.RetriesCount,
	// )

	pg, err := postgres.Connection(cfg.Postgres.Url)
	if err != nil {
		panic(err)
	}
	defer pg.Close()

	log.Info("Initializing repositories...")
	repository := repo.NewRepository(pg)

	log.Info("Initializing services...")
	dependencies := service.ServiceDependencies{Repo: repository}
	service := service.NewService(dependencies)

	handler := echo.New()
	handler.Use(logRequest)
	v1.NewRouter(handler, service)

	for _, route := range handler.Routes() {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	handler.Logger.Fatal(handler.Start(":8080"))
}

func logRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		fmt.Printf("Request - Method: %s, Path: %s\n", req.Method, req.URL.Path)
		return next(c)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
