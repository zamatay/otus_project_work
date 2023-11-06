package sysInfo

import (
	"bufio"
	"fmt"
	"io"
)

type ProcessorStat struct {
	cpu    string
	User   int64
	Nice   int64
	System int64
	Idle   int64
}

func GetProcessorStat() (*ProcessorStat, error) {
	return GetByPath[ProcessorStat]("/proc/stat", collectProcessor)
}

func collectProcessor(file io.Reader) (*ProcessorStat, error) {
	var result ProcessorStat
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
	fmt.Sscanf(string(line), "%s %d %d %d %d %f %f %f",
		&result.cpu, &result.User, &result.Nice, &result.System, &result.Idle)
	return &result, nil
}
