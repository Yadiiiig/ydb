package main

import (
	"fmt"
	"log"
	"net"

	"yadiiig.dev/ydb/internals/handler"
	reader "yadiiig.dev/ydb/internals/reader"
)

func main() {
	d := reader.ReadData("./idea.json", "./data.ydb")
	fmt.Println("Database is ready to Go.")
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler.NewGrpcServer(lis, d)
}
