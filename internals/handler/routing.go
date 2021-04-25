package handler

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "yadiiig.dev/ydb/internals/proto"
	reader "yadiiig.dev/ydb/internals/reader"
)

type grcpDetails struct {
	Server *grpc.Server
	reader.Drivers
}

func GrpcInit(lis net.Listener, d reader.Drivers) {
	grpcServer := grpc.NewServer()
	gd := grcpDetails{grpcServer, d}

	gd.selectQuery()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (g *grcpDetails) selectQuery() {
	service := NewSelectService(g.Drivers)
	pb.RegisterSelectServer(g.Server, service)
}
