package main

import (
	"github.com/wlcmtunknwndth/GoBotManager/internal/config"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
	envTest = "test"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Loaded config", slog.Any("config", cfg))

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {

	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envTest:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
