package handler

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "yadiiig.dev/ydb/internals/proto"
	reader "yadiiig.dev/ydb/internals/reader"
)

func NewGrpcServer(lis net.Listener, d *reader.Drivers) {
	s := grpc.NewServer()

	selectService := NewSelectService(d)
	insertService := NewInsertService(d)

	pb.RegisterSelectServer(s, selectService)
	pb.RegisterInsertServer(s, insertService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
