package storage

import (
	"project_work/internal/domain/models"
	"project_work/internal/services/grpc/monitor_v1"
	"reflect"
	"testing"
)

func Test_getDiskNode(t *testing.T) {
	type Args struct {
		sysMonitor models.SysMonitor
	}
	tests := []struct {
		name string
		args Args
		want []*monitor_v1.DiskInfoN
	}{
		{name: "test1", args: Args{
			sysMonitor: models.SysMonitor{Time: 0, Memory: models.Memory{}, AvgSysLoad: models.LoadAVG{}, AvgCpuLoad: models.CpuStat{}, DiskUsedFS: map[string]models.DiskInfoFS{},
				DiskUsedN: map[string]models.DiskInfoN{
					"tmpfs":          {Fs: "tmpfs", Node: 175564, Uses: 166, Free: 175398, UsePercent: "/run/user/1000", Mount: ""},
					"/dev/nvme0n1p5": {Fs: "/dev/nvme0n1p5", Node: 7045120, Uses: 1013733, Free: 6031387, UsePercent: "/", Mount: ""},
					"/dev/nvme0n1p1": {Fs: "/dev/nvme0n1p1", Node: 0, Uses: 0, Free: 0, UsePercent: "/boot/efi", Mount: ""},
				},
				DiskInfo: []models.DiskInfoAny{}, AvgCpu: models.AvgCPU{}, NetInfo: []models.NetInfo{}, StatInfo: &models.IOStat{}},
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

func Test_getNet(t *testing.T) {
	type Args struct {
		net []models.NetInfo
	}
	tests := []struct {
		name string
		args Args
		want *monitor_v1.NetStat
	}{
		{
			name: "Test_getNet",
			args: Args{net: []models.NetInfo{
				{State: "LISTEN",
					RecvQ:        10,
					SendQ:        20,
					LocalAddress: "0.0.0.0:ms-sql-s",
					PeerAddress:  "0.0.0.0:*",
				},
				{State: "LISTEN",
					RecvQ:        20,
					SendQ:        40,
					LocalAddress: "0.0.0.0:postgresql",
					PeerAddress:  "0.0.0.0:*",
				},
			},
			},
			want: &monitor_v1.NetStat{
				NetInfos: []*monitor_v1.NetInfo{
					{State: "LISTEN",
						RecvQ:        10,
						SendQ:        20,
						LocalAddress: "0.0.0.0:ms-sql-s",
						PeerAddress:  "0.0.0.0:*",
					},
					{State: "LISTEN",
						RecvQ:        20,
						SendQ:        40,
						LocalAddress: "0.0.0.0:postgresql",
						PeerAddress:  "0.0.0.0:*",
					},
				},
				StateInfos: []*monitor_v1.StateInfo{
					{State: "LISTEN", Count: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNet(tt.args.net); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiskRwPs(t *testing.T) {
	type Args struct {
		values models.SysMonitor
	}
	tests := []struct {
		name string
		args Args
		want []*monitor_v1.DiskUsedFS
	}{
		{
			name: "Test_getDiskRwPs",
			args: Args{models.SysMonitor{DiskUsedFS: map[string]models.DiskInfoFS{
				"one": {Used: "158", Available: "175406", UsedPercent: "1%"},
			}}},
			want: []*monitor_v1.DiskUsedFS{
				{Used: "158", Available: "175406", UsedPercent: "1%"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiskRwPs(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiskRwPs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiskTps(t *testing.T) {
	type Args struct {
		tps   []*monitor_v1.DiskTps
		value models.SysMonitor
	}
	tests := []struct {
		name string
		args Args
		want []*monitor_v1.DiskTps
	}{
		{
			name: "Test_getDiskTps",
			args: Args{
				tps: []*monitor_v1.DiskTps{
					{DiskDevice: "loop0", Tps: 0.01, KBReadS: 1165, KBWrtnS: 1165},
					{DiskDevice: "loop1", Tps: 0.22, KBReadS: 26322, KBWrtnS: 1165},
				},
			},
			want: []*monitor_v1.DiskTps{
				{DiskDevice: "loop0", Tps: 0.01, KBReadS: 1165, KBWrtnS: 1165},
				{DiskDevice: "loop1", Tps: 0.22, KBReadS: 26322, KBWrtnS: 1165},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiskTps(tt.args.tps, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiskTps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiskTpsAvg(t *testing.T) {
	type Args struct {
		tps   []*monitor_v1.DiskTps
		count float32
	}
	tests := []struct {
		name       string
		args       Args
		wantResult []*monitor_v1.DiskTps
	}{
		{name: "Test_getDiskTpsAvg", args: Args{
			tps: []*monitor_v1.DiskTps{
				{DiskDevice: "One", Tps: 4.0, KBReadS: 6.0, KBWrtnS: 10.0},
				{DiskDevice: "Two", Tps: 2.0, KBReadS: 4.0, KBWrtnS: 6.0},
			}, count: 2},
			wantResult: []*monitor_v1.DiskTps{
				{DiskDevice: "One", Tps: 2.0, KBReadS: 3.0, KBWrtnS: 5.0},
				{DiskDevice: "Two", Tps: 1.0, KBReadS: 2.0, KBWrtnS: 3.0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getDiskTpsAvg(tt.args.tps, tt.args.count); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("getDiskTpsAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_getDiscTpsStruct(t *testing.T) {
	type Args struct {
		DiskDevice string
		Tps        float32
		KBRead     float32
		KBWrtn     float32
	}
	tests := []struct {
		name string
		args Args
		want monitor_v1.DiskTps
	}{
		{
			name: "Test_getDiscTpsStruct",
			args: Args{DiskDevice: "Test", Tps: 1.0, KBRead: 10.0, KBWrtn: 11.0},
			want: monitor_v1.DiskTps{DiskDevice: "Test", Tps: 1.0, KBReadS: 10.0, KBWrtnS: 11.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiscTpsStruct(tt.args.DiskDevice, tt.args.Tps, tt.args.KBRead, tt.args.KBWrtn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiscTpsStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
