package main

import (
	"github.com/Sorokin41/test_task_go/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func main() {
	cfg := config.MustLoad()
	log := setLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("starting server", slog.String("address", cfg.Address))
	log.Debug("logger debug enabled")
}

func setLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
