package main

import (
	"testing"

	"github.com/google/uuid"
)

// var clients *Clients
// var ctx *Ctx
// var selectService *SelectService
// var insertService *InsertService
// var deleteService *DeleteService

// func init() {
// 	c := ConnectionSetup("localhost:8008")

// 	selectService = NewSelectService(c)
// 	insertService = NewInsertService(c)
// 	deleteService = NewDeleteService(c)

// 	ctx = ContextSetup()
// }

// func BenchmarkSelectSpec(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ctx.Delete(deleteService.Delete, "users", []*pb.DValues{{Row: "company", Operator: "=", Value: "Curabitur LLC"}})
// 		b.StopTimer()
// 		ctx.Insert(insertService.Insert, "users", []*pb.IValues{
// 			{Row: "userid", Value: "200i"},
// 			{Row: "firstname", Value: "Brenden"},
// 			{Row: "lastname", Value: "Baker"},
// 			{Row: "email", Value: "velit.Sed.malesuada@velitegestaslacinia.ca"},
// 			{Row: "company", Value: "Curabitur LLC"},
// 		}) //users: {"id": "99i", "firstname": "Brenden", "lastname": "Baker", "email": "velit.Sed.malesuada@velitegestaslacinia.ca", "company": "Curabitur LLC"}
// 		// fmt.Println(ctx.SelectSpec(selectService.Select, "users", []string{"*"}, []*pb.SValues{{Row: "company", Operator: "=", Value: "Curabitur LLC"}}))
// 		b.StartTimer()
// 	}
// }

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid.New()
		//fmt.Println(id.String())
	}
}

// func BenchmarkInsert(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ctx.Insert(insertService.Insert, "posts", []*pb.IValues{
// 			{Row: "userid", Value: "5"},
// 			{Row: "title", Value: "party"},
// 			{Row: "body", Value: "sick party"},
// 			{Row: "noexist", Value: "sick party"},
// 		})
// 	}
// }
