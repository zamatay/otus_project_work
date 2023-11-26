package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"project_work/internal/app"
	"project_work/internal/config"
	"project_work/internal/log"
	"project_work/internal/services/grpc"
	"syscall"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())

	cfg := config.Load()
	//
	log.SetupLogger(cfg.DevEnv)
	srvInfo := app.InitService()

	done := make(chan struct{})
	go shutDown(done)

	go grpc.Serve(cfg.Grpc, srvInfo)

	if err := app.ServiceRun(ctx, cancelFn); err != nil {
		log.Logger.Log.Error("Ошибка при старте сервиса", "Error", err)
		return
	}

	<-done
	cancelFn()
}

func shutDown(done chan struct{}) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Logger.Log.Info(fmt.Sprintf("Получен сигнал на завершение: %v.", <-quit))

	close(done)
}
