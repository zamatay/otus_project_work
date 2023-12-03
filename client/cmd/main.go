package main

import (
	"client_monitor/internal/app"
	"client_monitor/internal/config"
	"client_monitor/internal/log"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())

	cfg := config.Load()
	//
	log.SetupLogger(cfg.DevEnv)

	done := make(chan struct{})

	go shutDown(done)

	if err := app.ClientRun(ctx, cancelFn, cfg, done); err != nil {
		log.Logger.Log.Error("Ошибка при старте клиента", "Error", err)
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
