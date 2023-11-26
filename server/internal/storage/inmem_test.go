package storage

import (
	"project_work/internal/domain/models"
	"project_work/internal/services/grpc/monitor_v1"
	"reflect"
	"testing"
)

func Test_getDiskNode(t *testing.T) {
	type args struct {
		sysMonitor models.SysMonitor
	}
	tests := []struct {
		name string
		args args
		want []*monitor_v1.DiskInfoN
	}{
		{name: "test1", args: args{
			sysMonitor: struct {
				Time       int64
				Memory     models.Memory
				AvgSysLoad models.LoadAVG
				AvgCpuLoad models.CpuStat
				DiskUsedFS map[string]models.DiskInfoFS
				DiskUsedN  map[string]models.DiskInfoN
				DiskInfo   []models.DiskInfoAny
				AvgCpu     models.AvgCPU
				NetInfo    []models.NetInfo
				StatInfo   models.IOStat
			}{Time: 0, Memory: models.Memory{}, AvgSysLoad: models.LoadAVG{}, AvgCpuLoad: models.CpuStat{}, DiskUsedFS: map[string]models.DiskInfoFS{},
				DiskUsedN: map[string]models.DiskInfoN{
					"tmpfs":          {Fs: "tmpfs", Node: 175564, Uses: 166, Free: 175398, UsePercent: "/run/user/1000", Mount: ""},
					"/dev/nvme0n1p5": {Fs: "/dev/nvme0n1p5", Node: 7045120, Uses: 1013733, Free: 6031387, UsePercent: "/", Mount: ""},
					"/dev/nvme0n1p1": {Fs: "/dev/nvme0n1p1", Node: 0, Uses: 0, Free: 0, UsePercent: "/boot/efi", Mount: ""},
				},
				DiskInfo: []models.DiskInfoAny{}, AvgCpu: models.AvgCPU{}, NetInfo: []models.NetInfo{}, StatInfo: models.IOStat{}},
		}, want: []*monitor_v1.DiskInfoN{
			{Uses: "166", Free: "175398", UsePercent: "/run/user/1000"}, {Uses: "1013733", Free: "6031387", UsePercent: "/"}, {Uses: "0", Free: "0", UsePercent: "/boot/efi"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiskNode(tt.args.sysMonitor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiskNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
