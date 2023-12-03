package sysInfo

import (
	"context"
	"errors"
	"os/exec"
	sysInfo "project_work/internal/collector/linux"
	"project_work/internal/config"
	"project_work/internal/domain/models"
	"project_work/internal/log"
	"project_work/internal/storage"
	"time"
)

type SysInfoSrv struct {
	ResultWork storage.MonitoringInterface
}

func NewSysInfoSrv() *SysInfoSrv {
	return &SysInfoSrv{
		ResultWork: storage.NewMonitoring(),
	}
}

func (service *SysInfoSrv) Serve(ctx context.Context, cfg *config.Config) error {
	for {
		info, err := snapshotSysInfo(cfg)
		if err != nil {
			log.Logger.Log.Error("Ошибка при получении снапшота системы", "Error", err)
		}
		service.ResultWork.AddItem(info)
		time.Sleep(time.Second)
	}
}

func snapshotSysInfo(cfg *config.Config) (*models.SysMonitor, error) {
	var errorrProcess error = nil
	var result = models.SysMonitor{}

	//память
	if cfg.Monitor.Memory {
		if memory, err := sysInfo.GetMemory(); err == nil {
			result.Memory = *memory
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	//средняя загрузка
	if cfg.Monitor.LoadAvg {
		if loadAVG, err := sysInfo.GetLoadAVG(); err == nil {
			result.AvgSysLoad = *loadAVG
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	//процессор
	if cfg.Monitor.ProcessStat {
		if processorStat, err := sysInfo.GetProcessorStat(); err == nil {
			result.AvgCpuLoad = *processorStat
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	//диск 1
	if cfg.Monitor.Disk {
		if d, err := sysInfo.GetDiskInfoDev(); err == nil {
			result.DiskUsedFS = *d
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	//диск 3
	if cfg.Monitor.Disk {
		if d, err := sysInfo.GetDiskInfoSecondary(); err == nil {
			result.DiskUsedN = *d
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	//Общая информация о компе
	if cfg.Monitor.Iostat {
		if ioStat, err := sysInfo.GetIOStat(); err == nil {
			result.StatInfo = ioStat
			result.DiskInfo = result.StatInfo.Sysstat.Hosts[0].Statistics[0].Disk
			result.AvgCpu = result.StatInfo.Sysstat.Hosts[0].Statistics[0].AvgCPU
		} else {
			errorrProcess = errors.Join(errorrProcess, err)
		}
	}

	if info, err := sysInfo.GetNetInfo(); err != nil {
		errorrProcess = errors.Join(errorrProcess, err)
	} else {
		result.NetInfo = *info
	}

	if errors.Is(errorrProcess, exec.ErrNotFound) {
		log.Logger.Fatal(errorrProcess)
	}

	return &result, errorrProcess
}
