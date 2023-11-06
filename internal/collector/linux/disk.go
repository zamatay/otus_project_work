package sysInfo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type DiskInfo struct {
	ReadCount  uint64 `json:"readCount"`
	WriteCount uint64 `json:"writeCount"`
	Tps        uint64 `json:"tps"`
	Name       string `json:"name"`
	ReadBytes  uint64 `json:"ReadBytes"`
	WriteBytes uint64 `json:"WriteBytes"`
	ReadTime   uint64 `json:"ReadTime"`
	WriteTime  uint64 `json:"WriteTime"`
	IoTime     uint64 `json:"IoTime"`
}

type DiskInfo2 struct {
	fs          string
	kblock      string
	used        string
	available   string
	usedPercent string
}

type DiskInfo3 struct {
	Fs         string
	Node       int64
	Uses       int64
	Free       int64
	UsePercent string
	Mount      string
}

const (
	SectorSize = 512
)

func GetDiskInfoDev() (*map[string]DiskInfo2, error) {
	output := ExecuteCommand("df", "-H")
	return collectDisk2(bytes.NewBuffer(output))
}

func GetDiskInfo() (*map[string]DiskInfo, error) {
	return GetByPath[map[string]DiskInfo]("/proc/diskstats", collectDisk)
}

func GetDiskInfo3() (*map[string]DiskInfo3, error) {
	output := ExecuteCommand("df", "-i")
	return collectDisk3(bytes.NewBuffer(output))
}

func collectDisk(file io.Reader) (*map[string]DiskInfo, error) {
	result := make(map[string]DiskInfo, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 14 {
			continue
		}
		name := fields[2]

		// начинаем парсить
		reads, err := strconv.ParseUint((fields[3]), 10, 64)
		if err != nil {
			return &result, err
		}
		rbytes, err := strconv.ParseUint((fields[5]), 10, 64)
		if err != nil {
			return &result, err
		}
		rtime, err := strconv.ParseUint((fields[6]), 10, 64)
		if err != nil {
			return &result, err
		}
		writes, err := strconv.ParseUint((fields[7]), 10, 64)
		if err != nil {
			return &result, err
		}
		wbytes, err := strconv.ParseUint((fields[9]), 10, 64)
		if err != nil {
			return &result, err
		}
		wtime, err := strconv.ParseUint((fields[10]), 10, 64)
		if err != nil {
			return &result, err
		}
		tps, err := strconv.ParseUint((fields[11]), 10, 64)
		if err != nil {
			return &result, err
		}
		iotime, err := strconv.ParseUint((fields[12]), 10, 64)
		if err != nil {
			return &result, err
		}
		d := DiskInfo{
			ReadBytes:  rbytes * SectorSize,
			WriteBytes: wbytes * SectorSize,
			ReadCount:  reads,
			WriteCount: writes,
			ReadTime:   rtime,
			WriteTime:  wtime,
			Tps:        tps,
			IoTime:     iotime,
		}

		d.Name = name

		result[name] = d
	}
	return &result, nil
}

func collectDisk2(ior io.Reader) (*map[string]DiskInfo2, error) {
	result := make(map[string]DiskInfo2, 0)
	scanner := bufio.NewScanner(ior)
	isFirst := true
	for scanner.Scan() {
		if isFirst {
			isFirst = false
			continue
		}
		line := scanner.Text()
		if strings.HasPrefix(line, "tmpfs") {
			continue
		}
		di := DiskInfo2{}
		fmt.Println(line)
		if in, err := fmt.Sscanf(string(line), "%s %d %d %d %s",
			&di.fs, &di.kblock, &di.used, &di.available, &di.usedPercent); err == nil {
			log.Println(in)
			result[di.fs] = di
		}
	}
	return &result, nil
}

func collectDisk3(buffer *bytes.Buffer) (*map[string]DiskInfo3, error) {
	result := make(map[string]DiskInfo3, 0)
	scanner := bufio.NewScanner(buffer)
	isFirst := true
	for scanner.Scan() {
		if isFirst {
			isFirst = false
			continue
		}
		line := scanner.Text()
		di := DiskInfo3{}
		if _, err := fmt.Sscanf(line, "%s %d %d %d %s %s",
			&di.Fs, &di.Node, &di.Uses, &di.Free, &di.UsePercent, &di.UsePercent); err == nil {
			result[di.Fs] = di
		}
		scanner.Text()
	}
	return &result, nil
}
