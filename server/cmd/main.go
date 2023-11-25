package main

import (
	"context"
	"project_work/internal/app"
	"project_work/internal/config"
	"project_work/internal/log"
	"project_work/internal/services/grpc"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())

	cfg := config.Load()
	//
	log.SetupLogger(cfg.Env)
	srvInfo := app.InitService()

	go grpc.Serve(cfg.Grpc, srvInfo)

	if err := app.ServiceRun(ctx, cancelFn); err != nil {
		log.Logger.Log.Error("Ошибка при старте сервиса", "Error", err)
		return
	}
}
