package storage

import (
	"project_work/internal/domain/models"
	"time"
)

const initLen = 100

type MonitoringInterface interface {
	AddItem(item *models.SysMonitor)
}

type Monitoring struct {
	items     []models.SysMonitor
	itemsTime map[int64]*models.SysMonitor
}

func NewMonitoring() MonitoringInterface {
	return &Monitoring{
		items:     make([]models.SysMonitor, 0, initLen),
		itemsTime: make(map[int64]*models.SysMonitor, initLen),
	}
}

func (m *Monitoring) AddItem(item *models.SysMonitor) {
	item.Time = time.Now().Unix()
	m.items = append(m.items, *item)
	m.itemsTime[item.Time] = item
}
