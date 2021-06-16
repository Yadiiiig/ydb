package main

import (
	"log"
	"net"
	"os"

	"yadiiig.dev/ydb/internals/background"
	"yadiiig.dev/ydb/internals/handler"
	reader "yadiiig.dev/ydb/internals/reader"
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
