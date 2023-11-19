package models

type CpuStat struct {
	Cpu    string
	User   int64
	Nice   int64
	System int64
	Idle   int64
}
