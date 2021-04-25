package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "yadiiig.dev/ydb/internals/proto"
)

type Clients struct {
	Select pb.SelectClient
}

type Ctx struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func main() {
	connection := ConnectionSetup("localhost:8008")
	//defer connection.Close()

	clients := ClientSetup(connection)

	ctx := ContextSetup()
	//defer ctx.Cancel()

	testVar := []exampleData{}
	res, err := ctx.Select(clients.Select, "users", []string{"*"})
	if err := json.Unmarshal([]byte(res), &testVar); err != nil {
		panic(err)
	}
	fmt.Print(testVar)
	fmt.Println(err)
}

func ConnectionSetup(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func ClientSetup(c *grpc.ClientConn) *Clients {
	return &Clients{
		Select: pb.NewSelectClient(c),
	}

}

func ContextSetup() *Ctx {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return &Ctx{
		Context: ctx,
		Cancel:  cancel,
	}
}

type exampleData struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Company   string `json:"company"`
}

func (ctx Ctx) Select(ec pb.SelectClient, t string, f []string) (string, error) {
	r, err := ec.SelectQuery(ctx.Context, &pb.SelectValues{Table: t, Fields: f})
	return r.GetResult(), err
}

func (ctx Ctx) SelectSpec(ec pb.SelectClient, t string, f []string, v []*pb.Values) (string, error) {
	r, err := ec.SelectQuery(ctx.Context, &pb.SelectValues{Table: t, Fields: f, Values: v})
	return r.GetResult(), err
}
