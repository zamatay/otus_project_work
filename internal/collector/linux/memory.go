//go:build linux
// +build linux

package sysInfo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	Total, Used, Buffers, Cached, Free, Available, Active, Inactive, SwapTotal, SwapUsed,
	SwapCached, SwapFree, Mapped, Shmem, Slab, PageTables, Committed, VmallocUsed uint64
	MemAvailableEnabled bool
}

func GetMemory() (*Memory, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectMemoryStats(file)
}

func collectMemoryStats(out io.Reader) (*Memory, error) {
	scanner := bufio.NewScanner(out)
	var memory Memory
	memStats := map[string]*uint64{
		"MemTotal":     &memory.Total,
		"MemFree":      &memory.Free,
		"MemAvailable": &memory.Available,
		"Buffers":      &memory.Buffers,
		"Cached":       &memory.Cached,
		"Active":       &memory.Active,
		"Inactive":     &memory.Inactive,
		"SwapCached":   &memory.SwapCached,
		"SwapTotal":    &memory.SwapTotal,
		"SwapFree":     &memory.SwapFree,
		"Mapped":       &memory.Mapped,
		"Shmem":        &memory.Shmem,
		"Slab":         &memory.Slab,
		"PageTables":   &memory.PageTables,
		"Committed_AS": &memory.Committed,
		"VmallocUsed":  &memory.VmallocUsed,
	}
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		fld := line[:i]
		if ptr := memStats[fld]; ptr != nil {
			val := strings.TrimSpace(strings.TrimRight(line[i+1:], "kB"))
			if v, err := strconv.ParseUint(val, 10, 64); err == nil {
				*ptr = v * 1024
			}
			if fld == "MemAvailable" {
				memory.MemAvailableEnabled = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/meminfo: %s", err)
	}

	memory.SwapUsed = memory.SwapTotal - memory.SwapFree

	if memory.MemAvailableEnabled {
		memory.Used = memory.Total - memory.Available
	} else {
		memory.Used = memory.Total - memory.Free - memory.Buffers - memory.Cached
	}

	return &memory, nil
}
