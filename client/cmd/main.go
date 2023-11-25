package main

import (
	"client_monitor/internal/app"
	"client_monitor/internal/config"
	"client_monitor/internal/log"
	"context"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())

	cfg := config.Load()
	//
	log.SetupLogger(cfg.Env)

	if err := app.ClientRun(ctx, cancelFn, cfg); err != nil {
		log.Logger.Log.Error("Ошибка при старте клиента", "Error", err)
		return
	}
}
