package storage

import (
	"errors"
	"fmt"
	"project_work/internal/domain/models"
	"project_work/internal/services/grpc/monitor_v1"
	"strconv"
	"sync"
	"time"
)

const initLen = 100

type MonitoringInterface interface {
	AddItem(item *models.SysMonitor)
	GetAvgLastData(sec int32) (model *monitor_v1.AllResponse, err error)
}

type Monitoring struct {
	Items []models.SysMonitor
	mu    sync.Mutex
}

func NewMonitoring() MonitoringInterface {
	return &Monitoring{
		Items: make([]models.SysMonitor, 0, initLen),
	}
}

func (m *Monitoring) AddItem(item *models.SysMonitor) {
	item.Time = time.Now().Unix()
	m.Items = append(m.Items, *item)
}

func (m *Monitoring) GetAvgLastData(sec int32) (*monitor_v1.AllResponse, error) {
	model := &monitor_v1.AllResponse{}
	m.mu.Lock()
	defer m.mu.Unlock()
	if int32(len(m.Items)) < sec {
		return nil, errors.New("SmallTime")
	}
	fromIndex := int32(len(m.Items)) - sec
	data := m.Items[fromIndex:]
	value := models.SysMonitor{}
	for _, value = range data {
		model.SysLoadAverage += float32(value.AvgSysLoad.Minute)
		if model.CpuAverage == nil {
			model.CpuAverage = &monitor_v1.CpuInfo{}
		}
		model.CpuAverage.User += int32(value.AvgCpuLoad.User)
		model.CpuAverage.Idle += int32(value.AvgCpuLoad.Idle)
		model.CpuAverage.System += int32(value.AvgCpuLoad.System)
		model.DiskTps = getDiskTps(model.DiskTps, value)
	}
	model.DiskRwPs = getDiskRwPs(value)
	model.DiskInfoN = getDiskNode(value)
	count_fl := float32(len(data))
	count := int32(len(data))
	model.SysLoadAverage = model.SysLoadAverage / count_fl
	model.CpuAverage.User = model.CpuAverage.User / count
	model.CpuAverage.Idle = model.CpuAverage.Idle / count
	model.CpuAverage.System = model.CpuAverage.System / count
	model.DiskTps = getDiskTpsAvg(model.DiskTps, count_fl)
	model.Net = getNet(value.NetInfo)
	return model, nil
}

func getNet(net []models.NetInfo) *monitor_v1.NetStat {
	result := monitor_v1.NetStat{
		NetInfos:   make([]*monitor_v1.NetInfo, 0, 0),
		StateInfos: make([]*monitor_v1.StateInfo, 0, 0),
	}
	si := make(map[string]int, 0)
	for _, value := range net {
		result.NetInfos = append(result.NetInfos,
			&monitor_v1.NetInfo{
				State:        value.State,
				RecvQ:        int32(value.RecvQ),
				SendQ:        int32(value.SendQ),
				LocalAddress: value.LocalAddress,
				PeerAddress:  fmt.Sprintf("%v", value.PeerAddress),
			})
		si[value.State] += 1
	}
	for key, value := range si {
		result.StateInfos = append(result.StateInfos, &monitor_v1.StateInfo{State: key, Count: int32(value)})
	}
	return &result
}

func getDiskNode(sysMonitor models.SysMonitor) []*monitor_v1.DiskInfoN {
	result := make([]*monitor_v1.DiskInfoN, 0, 1)
	for _, value := range sysMonitor.DiskUsedN {
		result = append(result, &monitor_v1.DiskInfoN{Uses: strconv.FormatInt(value.Uses, 10), Free: strconv.FormatInt(value.Free, 10), UsePercent: value.UsePercent})
	}
	return result
}

func getDiskRwPs(values models.SysMonitor) []*monitor_v1.DiskUsedFS {
	result := make([]*monitor_v1.DiskUsedFS, 0, 1)
	for _, value := range values.DiskUsedFS {
		result = append(result, &monitor_v1.DiskUsedFS{Used: value.Used, Available: value.Available, UsedPercent: value.UsedPercent})
	}
	return result
}

func getDiskTps(tps []*monitor_v1.DiskTps, value models.SysMonitor) []*monitor_v1.DiskTps {
	if tps == nil {
		tps = make([]*monitor_v1.DiskTps, 0, 1)
	}
	dis := []*monitor_v1.DiskTps{}
	var diskTpsAll = make(map[string]monitor_v1.DiskTps, len(tps))
	var diskTpsCurrent = make(map[string]monitor_v1.DiskTps, len(tps))
	for _, v := range tps {
		diskTpsAll[v.DiskDevice] = *v
	}
	for _, v := range value.DiskInfo {
		if v.Tps > 0 || v.KBRead > 0 || v.KBWrtn > 0 {
			diskTpsCurrent[v.DiskDevice] = getDiscTpsStruct(v.DiskDevice, float32(v.Tps), float32(v.KBRead), float32(v.KBWrtn))
		}
	}
	for key, value := range diskTpsCurrent {
		diskTpsAll[key] = getDiscTpsStruct(diskTpsAll[key].DiskDevice, diskTpsAll[key].Tps+value.Tps, diskTpsAll[key].KBReadS+value.KBReadS, diskTpsAll[key].KBWrtnS+value.KBWrtnS)
	}

	for key, d := range diskTpsAll {
		t := getDiscTpsStruct(key, float32(d.Tps), float32(d.KBReadS), float32(d.KBWrtnS))
		dis = append(dis, &t)
	}
	return dis
}

func getDiskTpsAvg(tps []*monitor_v1.DiskTps, count float32) (result []*monitor_v1.DiskTps) {
	for _, value := range tps {
		result = append(result, &monitor_v1.DiskTps{
			DiskDevice: value.DiskDevice,
			Tps:        value.Tps / count,
			KBWrtnS:    value.KBWrtnS / count,
			KBReadS:    value.KBReadS / count,
		})
	}
	return result
}

func getDiscTpsStruct(DiskDevice string, Tps float32, KBRead float32, KBWrtn float32) monitor_v1.DiskTps {
	return monitor_v1.DiskTps{DiskDevice: DiskDevice, Tps: Tps, KBReadS: KBRead, KBWrtnS: KBWrtn}
}
