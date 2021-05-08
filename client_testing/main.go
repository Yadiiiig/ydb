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

type Services struct {
	Select pb.SelectClient
	Insert pb.InsertClient
}

type SelectService struct {
	Select pb.SelectClient
}

type InsertService struct {
	Insert pb.InsertClient
}

type Ctx struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func main() {
	c := ConnectionSetup("localhost:8008")
	defer c.Close()

	selectService := NewSelectService(c)
	insertService := NewInserService(c)

	ctx := ContextSetup()
	defer ctx.Cancel()

	testVar := []exampleData{}
	// res, err := ctx.Select(clients.Select, "users", []string{"*"})
	resQ, _ := ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.Values{{Operator: "=", Row: "firstname", Value: "Piper"}})
	if err := json.Unmarshal([]byte(resQ), &testVar); err != nil {
		panic(err)
	}

	resI, _ := ctx.Insert(insertService.Insert, "posts", []*pb.IValues{
		{Row: "userid", Value: "5"},
		{Row: "title", Value: "party"},
		{Row: "body", Value: "sick party"},
		{Row: "noexist", Value: "sick party"},
	})
	fmt.Println(resI)

	fmt.Print(testVar)
}

func ConnectionSetup(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func NewSelectService(c *grpc.ClientConn) *SelectService {
	return &SelectService{
		Select: pb.NewSelectClient(c),
	}

}

func NewInserService(c *grpc.ClientConn) *InsertService {
	return &InsertService{
		Insert: pb.NewInsertClient(c),
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

func (ctx Ctx) Insert(ec pb.InsertClient, t string, v []*pb.IValues) (bool, error) {
	r, err := ec.InsertQuery(ctx.Context, &pb.InsertValues{Table: t, Values: v})
	return r.GetResult(), err
}
