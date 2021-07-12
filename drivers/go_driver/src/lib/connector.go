package lib

import (
	"context"
	"time"

	pb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib/proto"
	"google.golang.org/grpc"
)

type Connection struct {
	Conn     *grpc.ClientConn
	Services *Services
	//Ctx      *Ctx
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

func Connect(address string) (*Connection, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &Connection{
		Conn:     conn,
		Services: ServiceSetup(conn),
		// Ctx:      ContextSetup(),
	}, nil
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return &Ctx{
		Context: ctx,
		Cancel:  cancel,
	}
}
