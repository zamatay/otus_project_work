package main

import (
	"fmt"
	sysInfo "project_work/internal/collector/linux"
)

func main() {
	//cfg := config.Load()
	//
	//initialization.SetupLogger(cfg.Env)

	memory, _ := sysInfo.GetMemory()
	fmt.Printf("%t, %v\n", memory, memory)

	loadAVG, _ := sysInfo.GetLoadAVG()
	fmt.Printf("%v\n", loadAVG)

	processorStat, _ := sysInfo.GetProcessorStat()
	fmt.Printf("%v\n", processorStat)

	//diskInfo, _ := sysInfo.GetDiskInfo()
	//fmt.Printf("%v\n", diskInfo)

	sysInfo.GetDiskInfoDev()

	ioStat := sysInfo.GetIOStat()

	diskInfo := ioStat.Sysstat.Hosts[0].Statistics[0].Disk
	cpuAvg := ioStat.Sysstat.Hosts[0].Statistics[0].AvgCPU

	fmt.Printf("%v\n%v\n", diskInfo, cpuAvg)

	diskInfo3, _ := sysInfo.GetDiskInfo3()
	fmt.Printf("%v/n", diskInfo3)
}
