package main

import (
	"context"
	"project_work/internal/app"
	"project_work/internal/config"
	"project_work/internal/log"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())

	cfg := config.Load()
	//
	log.SetupLogger(cfg.Env)
	app.InitService()

	if err := app.ServiceRun(ctx, cancelFn); err != nil {
		log.Logger.Log.Error("Ошибка при старте сервиса", "Error", err)
		return
	}
}
