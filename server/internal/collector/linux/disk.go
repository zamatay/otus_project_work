package sysInfo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"project_work/internal/domain/models"
	"project_work/internal/log"
	"strconv"
	"strings"
)

const (
	SectorSize = 512
)

func GetDiskInfoDev() (*map[string]models.DiskInfoFS, error) {
	output, err := ExecuteCommand("df", "-H")
	if err != nil {
		log.Logger.Log.Error("Ошибка при получении информации о дисках. Необходимо установить утилиту df", err)
		return nil, err
	}
	return collectDisk2(bytes.NewBuffer(output)), nil
}

//func GetDiskInfo() (*map[string]models.DiskInfo, error) {
//	return GetByPath[map[string]models.DiskInfo]("/proc/diskstats", collectDisk)
//}

func GetDiskInfoSecondary() (*map[string]models.DiskInfoN, error) {
	output, err := ExecuteCommand("df", "-i")
	if err != nil {
		log.Logger.Log.Error("Ошибка при получении вторичной информации о дисках. Необходимо установить утилуту df", err)
		return nil, err
	}
	return collectDisk3(bytes.NewBuffer(output))
}

func collectDisk(file io.Reader) (*map[string]models.DiskInfo, error) {
	result := make(map[string]models.DiskInfo, 0)
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
		d := models.DiskInfo{
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

func collectDisk2(ior io.Reader) *map[string]models.DiskInfoFS {
	result := make(map[string]models.DiskInfoFS, 0)
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
		di := models.DiskInfoFS{}
		if _, err := fmt.Sscanf(string(line), "%s %s %s %s %s",
			&di.Fs, &di.Kblock, &di.Used, &di.Available, &di.UsedPercent); err == nil {
			result[di.Fs] = di
		}
	}
	return &result
}

func collectDisk3(buffer *bytes.Buffer) (*map[string]models.DiskInfoN, error) {
	result := make(map[string]models.DiskInfoN, 0)
	scanner := bufio.NewScanner(buffer)
	isFirst := true
	for scanner.Scan() {
		if isFirst {
			isFirst = false
			continue
		}
		line := scanner.Text()
		di := models.DiskInfoN{}
		if _, err := fmt.Sscanf(line, "%s %d %d %d %s %s",
			&di.Fs, &di.Node, &di.Uses, &di.Free, &di.UsePercent, &di.Mount); err == nil {
			result[di.Fs] = di
		}
		scanner.Text()
	}
	return &result, nil
}
