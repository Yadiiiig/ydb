package main

import (
	"testing"

	pb "yadiiig.dev/ydb/internals/proto"
)

var clients *Clients
var ctx *Ctx

func init() {
	connection := ConnectionSetup("localhost:8008")
	//defer connection.Close()

	clients = ClientSetup(connection)

	ctx = ContextSetup()
	//defer ctx.Cancel()
}

// func BenchmarkSelect(b *testing.B) {
// 	ctx.Select(clients.Select, "users", []string{"*"})
// }

func BenchmarkSelectSpec(b *testing.B) {
	ctx.SelectSpec(clients.Select, "users", []string{"*"}, []*pb.Values{{Operator: "=", Row: "firstname", Value: "Piper"}})
}
