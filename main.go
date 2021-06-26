package main

import (
	"log"
	"net"
	"os"

	"github.com/Yadiiiig/ydb/internals/background"
	"github.com/Yadiiiig/ydb/internals/handler"
	reader "github.com/Yadiiiig/ydb/internals/reader"
)

func main() {
	d, err := reader.ReadData("data/")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	background.ExitHandler(d)
	background.BackgroundUpdating(d)
	log.Println("Database is ready to Go.")
	handler.NewGrpcServer(lis, d)
}
