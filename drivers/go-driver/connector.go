package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "yadiiig.dev/ydb/go-driver/proto"
)

type Connection struct {
	Conn     *grpc.ClientConn
	Services *Services
	Ctx      *Ctx
}

type Services struct {
	selectService pb.SelectClient
	insertService pb.InsertClient
	deleteService pb.DeleteClient
	updateService pb.UpdateClient
}

type Ctx struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func connect(address string) *Connection {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &Connection{
		Conn:     conn,
		Services: serviceSetup(conn),
		Ctx:      contextSetup(),
	}
}

func serviceSetup(c *grpc.ClientConn) *Services {
	return &Services{
		selectService: pb.NewSelectClient(c),
		insertService: pb.NewInsertClient(c),
		deleteService: pb.NewDeleteClient(c),
		updateService: pb.NewUpdateClient(c),
	}
}

func contextSetup() *Ctx {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return &Ctx{
		Context: ctx,
		Cancel:  cancel,
	}
}
