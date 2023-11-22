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
		Host:     "127.0.0.1",
		Port:     "4001",
		User:     "",
		Password: "",
		Database: "public",
	}
	err := greptimedb.Setup()
	if err != nil {
		panic(err)
	}

	pg = &db.Postgres{
		Host:     "127.0.0.1",
		Port:     "4003",
		User:     "",
		Password: "",
		Database: "public",
	}
	err = pg.Setup()
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := greptimedb.Insert(); err != nil {
		panic(err)
	} else {
		fmt.Println("insert success")
	}

	if all, err := pg.AllMonitors(); err != nil {
		panic(err)
	} else {
		fmt.Println(all)
	}
}
