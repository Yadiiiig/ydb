package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "yadiiig.dev/ydb/internals/proto"
)

type Services struct {
	Select pb.SelectClient
	Insert pb.InsertClient
	Delete pb.DeleteClient
	Update pb.UpdateClient
}

type SelectService struct {
	Select pb.SelectClient
}

type InsertService struct {
	Insert pb.InsertClient
}

type DeleteService struct {
	Delete pb.DeleteClient
}

type UpdateService struct {
	Update pb.UpdateClient
}

type Ctx struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func main() {
	c := ConnectionSetup("localhost:8008")
	defer c.Close()

	// selectService := NewSelectService(c)
	// insertService := NewInsertService(c)
	// deleteService := NewDeleteService(c)
	updateService := NewUpdateService(c)

	ctx := ContextSetup()
	defer ctx.Cancel()

	// fmt.Println("Inserting a user: Yadiiiig Yadiiiigson yadiiiig.dev@gmail.com No company")
	// ctx.Insert(insertService.Insert, "users", []*pb.IValues{
	// 	{Row: "userid", Value: "101i"},
	// 	{Row: "firstname", Value: "Yadiiiig"},
	// 	{Row: "lastname", Value: "Yadiiiigson"},
	// 	{Row: "email", Value: "yadiiiig.dev@gmail.com"},
	// 	{Row: "company", Value: "No company"},
	// })
	// // fmt.Println(x, err)
	// fmt.Println("Getting a this yadiiiig user using username = yadiiiig")
	// resQ, _ := ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.SValues{{Operator: "=", Row: "firstname", Value: "Yadiiiig"}})
	// fmt.Println(resQ)

	// fmt.Println("Updating yadiiiig's company to yadiiiig's inc. (using firstname and email")
	ctx.Update(updateService.Update, "users",
		[]*pb.MatchValues{
			{Row: "firstname", Operator: "=", Value: "Yadiiiig"},
			{Row: "email", Operator: "=", Value: "yadiiiig.dev@gmail.com"},
		},
		[]*pb.UValues{
			{Row: "company", Value: "booyaa."},
		})
	// fmt.Println("Deleting this user")
	// ctx.Delete(deleteService.Delete, "users", []*pb.DValues{{Row: "firstname", Operator: "=", Value: "Yadiiiig"}, {Row: "email", Operator: "=", Value: "yadiiiig.dev@gmail.com"}})

	// fmt.Println("Trying to get this user again")
	// resQq, _ := ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.SValues{{Operator: "=", Row: "firstname", Value: "Yadiiiig"}})
	// fmt.Println(resQq)
	//fmt.Println(testUpdate, err)
	// testVar := []exampleData{}
	// // res, err := ctx.Select(clients.Select, "users", []string{"*"})
	// if err := json.Unmarshal([]byte(resQ), &testVar); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resI)

	// fmt.Print(testVar)
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

func NewInsertService(c *grpc.ClientConn) *InsertService {
	return &InsertService{
		Insert: pb.NewInsertClient(c),
	}
}

func NewDeleteService(c *grpc.ClientConn) *DeleteService {
	return &DeleteService{
		Delete: pb.NewDeleteClient(c),
	}
}

func NewUpdateService(c *grpc.ClientConn) *UpdateService {
	return &UpdateService{
		Update: pb.NewUpdateClient(c),
	}
}

func ContextSetup() *Ctx {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

func (ctx Ctx) SelectSpec(ec pb.SelectClient, t string, f []string, v []*pb.SValues) (string, error) {
	r, err := ec.SelectQuery(ctx.Context, &pb.SelectValues{Table: t, Fields: f, Values: v})
	return r.GetResult(), err
}

func (ctx Ctx) Insert(ec pb.InsertClient, t string, v []*pb.IValues) (bool, error) {
	r, err := ec.InsertQuery(ctx.Context, &pb.InsertValues{Table: t, Values: v})
	return r.GetResult(), err
}

func (ctx Ctx) Delete(ec pb.DeleteClient, t string, v []*pb.DValues) (bool, error) {
	r, err := ec.DeleteQuery(ctx.Context, &pb.DeleteValues{Table: t, Values: v})
	return r.GetResult(), err
}

func (ctx Ctx) Update(ec pb.UpdateClient, t string, v []*pb.MatchValues, d []*pb.UValues) (bool, error) {
	r, err := ec.UpdateQuery(ctx.Context, &pb.UpdateValues{Table: t, Matchers: v, Values: d})
	return r.GetResult(), err
}
