package handler

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	reader "github.com/Yadiiiig/ydb/internals/reader"
)

func NewGrpcServer(lis net.Listener, d *reader.Drivers) {
	s := grpc.NewServer()

	selectService := NewSelectService(d)
	insertService := NewInsertService(d)
	deleteService := NewDeleteService(d)
	updateService := NewUpdateService(d)

	pb.RegisterSelectServer(s, selectService)
	pb.RegisterInsertServer(s, insertService)
	pb.RegisterDeleteServer(s, deleteService)
	pb.RegisterUpdateServer(s, updateService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
