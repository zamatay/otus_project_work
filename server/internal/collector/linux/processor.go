package sysInfo

import (
	"bufio"
	"fmt"
	"io"
	"project_work/internal/domain/models"
)

func GetProcessorStat() (*models.CpuStat, error) {
	return GetByPath[models.CpuStat]("/proc/stat", collectProcessor)
}

func collectProcessor(file io.Reader) (*models.CpuStat, error) {
	var result models.CpuStat
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(string(line), "%s %d %d %d %d %f %f %f",
		&result.Cpu, &result.User, &result.Nice, &result.System, &result.Idle)
	return &result, nil
}
