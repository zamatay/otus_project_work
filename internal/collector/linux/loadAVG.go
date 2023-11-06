package sysInfo

import (
	"bufio"
	"fmt"
	"io"
)

type LoadAVG struct {
	minute   float64
	minute5  float64
	Minute10 float64
}

func GetLoadAVG() (*LoadAVG, error) {
	return GetByPath[LoadAVG]("/proc/loadavg", collectLoadAvg)
}

func collectLoadAvg(file io.Reader) (*LoadAVG, error) {
	var result LoadAVG
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(string(line), "%f %f %f %d/%d %d",
			&result.minute, &result.minute5, &result.Minute10)
	}
	return &result, nil
}
