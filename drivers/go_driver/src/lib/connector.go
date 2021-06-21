package lib

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "yadiiig.dev/ydb/go_driver/src/lib/proto"
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

func Connect(address string) *Connection {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &Connection{
		Conn:     conn,
		Services: ServiceSetup(conn),
		Ctx:      ContextSetup(),
	}
}

func ServiceSetup(c *grpc.ClientConn) *Services {
	return &Services{
		selectService: pb.NewSelectClient(c),
		insertService: pb.NewInsertClient(c),
		deleteService: pb.NewDeleteClient(c),
		updateService: pb.NewUpdateClient(c),
	}
}

func ContextSetup() *Ctx {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return &Ctx{
		Context: ctx,
		Cancel:  cancel,
	}
}
