package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"project_work/internal/log"
	"project_work/internal/services/sysInfo"
	"syscall"
)

var (
	serviceSysInfo *sysInfo.SysInfoSrv
)

func InitService() {
	serviceSysInfo = sysInfo.NewSysInfoSrv()
}

func ServiceRun(ctx context.Context, cancelFunc context.CancelFunc) error {
	done := make(chan struct{})

	go shutDown(done)

	go serviceSysInfo.Serve(ctx)

	<-done
	cancelFunc()
	return nil

}

func shutDown(done chan struct{}) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Logger.Log.Info(fmt.Sprintf("Получен сигнал на завершение: %v.", <-quit))

	close(done)
}
