package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

	DB *gorm.DB
}

// Setup is to init the DB, and SHOULD BE called only once
func (p *Postgres) Setup() error {
	if p.DB != nil {
		return nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		p.Host, p.User, p.Password, p.Database, p.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = db
	return nil
}
