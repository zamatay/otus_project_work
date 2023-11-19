package sysInfo

import (
	"bufio"
	"fmt"
	"io"
	"project_work/internal/domain/models"
)

func GetLoadAVG() (*models.LoadAVG, error) {
	return GetByPath[models.LoadAVG]("/proc/loadavg", collectLoadAvg)
}

func collectLoadAvg(file io.Reader) (*models.LoadAVG, error) {
	var result models.LoadAVG
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(string(line), "%f %f %f %d/%d %d",
			&result.Minute, &result.Minute5, &result.Minute10)
	}
	return &result, nil
}
