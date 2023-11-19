package sysInfo

import (
	"bufio"
	"bytes"
	"io"
	"project_work/internal/domain/models"
	"strconv"
	"strings"
)

func GetNetInfo() (*[]models.NetInfo, error) {
	output := ExecuteCommand("ss", "-ta")
	return collectNetInfo(bytes.NewBuffer(output))
}

func collectNetInfo(io io.Reader) (*[]models.NetInfo, error) {
	result := make([]models.NetInfo, 0, 10)
	scanner := bufio.NewScanner(io)
	isFirst := true
	for scanner.Scan() {
		if isFirst {
			isFirst = false
			continue
		}
		line := scanner.Text()
		ni := models.NetInfo{}
		fields := strings.Fields(line)
		ni.State = fields[0]
		ni.RecvQ, _ = strconv.Atoi(fields[1])
		ni.SendQ, _ = strconv.Atoi(fields[2])
		ni.LocalAddress = fields[3]
		ni.PeerAddress = fields[4]
		result = append(result, ni)
	}
	return &result, nil
}
