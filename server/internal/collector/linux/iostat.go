package sysInfo

import (
	"encoding/json"
	"project_work/internal/domain/models"
)

func GetIOStat() (*models.IOStat, error) {
	data, err := ExecuteCommand("iostat", "-o", "JSON")
	if err != nil {
		return nil, err
	}
	value := models.IOStat{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return &value, nil
}
