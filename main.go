package main

import (
	"fmt"

	"github.com/greptimeteam/gorm-greptimedb/db"
)

var (
	greptimedb *db.Greptime
	pg         *db.Postgres
)

func init() {
	greptimedb = &db.Greptime{
		Host: "127.0.0.1",
		Port: "4001",
	}
	greptimedb.Setup()

	pg = &db.Postgres{
		Host:     "127.0.0.1",
		Port:     "4003",
		User:     "",
		Password: "",
		Database: "public",
	}
	pg.Setup()
}

func main() {
	fmt.Println("hello")
}
