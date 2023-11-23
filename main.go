package main

import (
	"fmt"

	"github.com/greptimeteam/gorm-greptimedb/db"
)

var (
	greptimedb *db.Greptime
	pg         *db.Postgres
	mysql      *db.Mysql
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

	mysql = &db.Mysql{
		Host:     "127.0.0.1",
		Port:     "4002",
		User:     "",
		Password: "",
		Database: "public",
	}
	err = mysql.Setup()
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := greptimedb.Insert(); err != nil {
		fmt.Printf("error of inserting into greptimedb: %v\n", err)
	} else {
		fmt.Println("insert success via greptimedb-client")
	}

	if all, err := mysql.AllMonitors(); err != nil {
		fmt.Printf("error of retrieving monitors via MySQL: %v\n", err)
	} else {
		fmt.Println("retrieving success via MySQL:")
		for i, m := range all {
			fmt.Printf("%d: %#v\n", i, m)
		}

	}

	if all, err := pg.AllMonitors(); err != nil {
		fmt.Printf("error of retrieving monitors via PostgresQL: %v\n", err)
	} else {
		fmt.Println("retrieving success via Postgres:")
		for i, m := range all {
			fmt.Printf("%d: %#v\n", i, m)
		}
	}
}
