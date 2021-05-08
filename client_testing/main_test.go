package main

import (
	"fmt"
	"testing"

	pb "yadiiig.dev/ydb/internals/proto"
)

// var clients *Clients
var ctx *Ctx
var selectService *SelectService
var insertService *InsertService

func init() {
	c := ConnectionSetup("localhost:8008")
	//defer c.Close()

	selectService = NewSelectService(c)
	insertService = NewInserService(c)

	ctx = ContextSetup()
	//fmt.Println("do i even execute?")
}

// func BenchmarkSelect(b *testing.B) {
// 	ctx.Select(clients.Select, "users", []string{"*"})
// }

func BenchmarkSelectSpec(b *testing.B) {
	r, err := ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.Values{{Operator: "=", Row: "firstname", Value: "Piper"}})
	fmt.Println(r, err)
	//ctx.Select(clients.Select, "users", []string{"*"})
}

func BenchmarkInsert(b *testing.B) {
	ctx.Insert(insertService.Insert, "posts", []*pb.IValues{
		{Row: "userid", Value: "5"},
		{Row: "title", Value: "party"},
		{Row: "body", Value: "sick party"},
		{Row: "noexist", Value: "sick party"},
	})
}
