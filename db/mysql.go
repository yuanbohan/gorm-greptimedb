package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

	DB *gorm.DB
}

// Setup is to init the DB, and SHOULD BE called only once
func (m *Mysql) Setup() error {
	if m.DB != nil {
		return nil
	}

	dsn := fmt.Sprintf("tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Host, m.Port, m.Database)
	if m.User != "" && m.Password != "" {
		dsn = fmt.Sprintf("%s:%s@%s", m.User, m.Password, dsn)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	m.DB = db
	return nil
}

func (p *Mysql) AllMonitors() ([]Monitor, error) {
	var monitors []Monitor
	err := p.DB.Find(&monitors).Error
	return monitors, err
}
