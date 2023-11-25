package storage

import (
	"errors"
	"project_work/internal/domain/models"
	"sync"
	"time"
)

const initLen = 100

type MonitoringInterface interface {
	AddItem(item *models.SysMonitor)
}

type Monitoring struct {
	Items []models.SysMonitor
	mu    sync.Mutex
}

func NewMonitoring() MonitoringInterface {
	return &Monitoring{
		Items: make([]models.SysMonitor, 0, initLen),
	}
}

func (m *Monitoring) AddItem(item *models.SysMonitor) {
	item.Time = time.Now().Unix()
	m.Items = append(m.Items, *item)
}

func (m *Monitoring) GetAvgLastData(sec time.Duration) (model *models.SysMonitor, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if time.Duration(len(m.Items)) < sec {
		return nil, errors.New("SmallTime")
	}
	fromIndex := time.Duration(len(m.Items)) - sec
	data := m.Items[fromIndex:]
	for key, value := range data {
		model.StatInfo := value.StatInfo
	}
}
