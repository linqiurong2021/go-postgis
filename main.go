package main

import (
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/linqiurong2021/go-postgis/conf"
	"github.com/linqiurong2021/go-postgis/server"
)

// Pool Pool
var Pool *pgxpool.Pool

func main() {
	//
	err := conf.InitConfig("./conf/conf.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	server := server.NewServer()
	server.Run()
}
