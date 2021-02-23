package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/linqiurong2021/go-postgis/conf"
	"github.com/linqiurong2021/go-postgis/libs"
)

func main() {
	//
	err := conf.InitConfig("./conf/conf.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	postgre := new(libs.Postgre)
	dbpool, err := postgre.Connect()
	if err != nil {
		fmt.Printf("connect postgre err:%s", err)
		os.Exit(1)
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	fmt.Println("Main")
}
