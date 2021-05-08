package main

import (
	"testing"

	pb "yadiiig.dev/ydb/internals/proto"
)

// var clients *Clients
var ctx *Ctx
var selectService *SelectService
var insertService *InsertService

func init() {
	c := ConnectionSetup("localhost:8008")

	selectService = NewSelectService(c)
	insertService = NewInserService(c)

	ctx = ContextSetup()
}

func BenchmarkSelectSpec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.Values{{Operator: "=", Row: "firstname", Value: "Piper"}})
	}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx.Insert(insertService.Insert, "posts", []*pb.IValues{
			{Row: "userid", Value: "5"},
			{Row: "title", Value: "party"},
			{Row: "body", Value: "sick party"},
			{Row: "noexist", Value: "sick party"},
		})
	}
}
