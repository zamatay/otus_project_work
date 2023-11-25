package sysInfo

import (
	"context"
	"errors"
	sysInfo "project_work/internal/collector/linux"
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

func (service *SysInfoSrv) Serve(ctx context.Context) error {
	for {
		info, err := snapshotSysInfo()
		if err != nil {
			log.Logger.Log.Error("Ошибка при получении снапшота системы", "Error", err)
		}
		service.ResultWork.AddItem(info)
		time.Sleep(time.Second)
	}
}

func snapshotSysInfo() (*models.SysMonitor, error) {
	var errorrProcess error = nil
	var result = models.SysMonitor{}

	//память
	if memory, err := sysInfo.GetMemory(); err == nil {
		result.Memory = *memory
	} else {
		errors.Join(errorrProcess, err)
	}

	//средняя загрузка
	if loadAVG, err := sysInfo.GetLoadAVG(); err == nil {
		result.AvgSysLoad = *loadAVG
	} else {
		errors.Join(errorrProcess, err)
	}

	//процессор
	if processorStat, err := sysInfo.GetProcessorStat(); err == nil {
		result.AvgCpuLoad = *processorStat
	} else {
		errors.Join(errorrProcess, err)
	}

	//диск 1
	if d, err := sysInfo.GetDiskInfoDev(); err == nil {
		result.DiskUsedFS = *d
	} else {
		errors.Join(errorrProcess, err)
	}

	//диск 2
	//if d, err := sysInfo.GetDiskInfo(); err == nil {
	//	result.DiskUsed = *d
	//} else {
	//	errors.Join(errorrProcess, err)
	//}

	//диск 3
	if d, err := sysInfo.GetDiskInfo3(); err == nil {
		result.DiskUsedN = *d
	} else {
		errors.Join(errorrProcess, err)
	}

	//Общая информация о компе
	if ioStat, err := sysInfo.GetIOStat(); err == nil {
		result.StatInfo = *ioStat
	} else {
		errors.Join(errorrProcess, err)
	}

	result.DiskInfo = result.StatInfo.Sysstat.Hosts[0].Statistics[0].Disk
	result.AvgCpu = result.StatInfo.Sysstat.Hosts[0].Statistics[0].AvgCPU

	if info, err := sysInfo.GetNetInfo(); err == nil {
		errors.Join(errorrProcess, err)
	} else {
		result.NetInfo = *info
	}

	return &result, errorrProcess
}
