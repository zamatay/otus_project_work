package app

import (
	"context"
	"project_work/internal/config"
	"project_work/internal/services/sysInfo"
)

var (
	serviceSysInfo *sysInfo.SysInfoSrv
)

func InitService() *sysInfo.SysInfoSrv {
	serviceSysInfo = sysInfo.NewSysInfoSrv()
	return serviceSysInfo
}

func ServiceRun(ctx context.Context, cancelFunc context.CancelFunc, cfg *config.Config) error {
	go serviceSysInfo.Serve(ctx, cfg)
	return nil
}
