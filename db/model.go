package db

import "time"

type Monitor struct {
	ID          int64     `gorm:"primaryKey"`
	Host        string    `gorm:"column:host"`
	Memory      uint64    `gorm:"column:memory"`
	Cpu         float64   `gorm:"column:cpu"`
	Temperature int64     `gorm:"column:temperature"`
	Ts          time.Time `gorm:"column:ts"`
}

func (Monitor) TableName() string {
	return "monitor"
}
