package log

import (
	log2 "log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Log struct {
	Log *slog.Logger
}

var (
	Logger Log
)

func SetupLogger(env string) {
	switch env {
	case envLocal:
		Logger.Log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		Logger.Log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		Logger.Log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
}

func (l *Log) Fatal(err error) {
	log2.Fatal(err)
}
