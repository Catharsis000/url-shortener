package main

import (
	"os"

	"golang.org/x/exp/slog"

	"github.com/Catharsis000/url-shortener.git/internal/config"
	"github.com/Catharsis000/url-shortener.git/internal/lib/logger/sl"
	"github.com/Catharsis000/url-shortener.git/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	//TODO: init config: cleanenv библиотека (add cleanenv with hithub)
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
}

// TODO: init logger: slog библиотека (import "log/slog")
func setupLogger(env string) *slog.Logger { // смотрим логи по конфигу
	var log *slog.Logger
	switch env {

	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
	//TODO: init storage: sqlite

	//TODO: init router: chi, "chi render" ("net/http")

	//TODO: run server
}
