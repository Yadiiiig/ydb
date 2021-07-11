package main

import (
	"fmt"
	"testing"

	ydb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib"
)

var db *ydb.Connection

// var bmData BenchmarkStruct
// var bmArray []BenchmarkStruct
var tmp []BenchmarkStruct

type BenchmarkStruct struct {
	ID        string
	Name      string
	Iteration string
}

// func BenchmarkInsert(b *testing.B) {
// 	db, _ = ydb.Connect("127.0.0.1:8008")
// 	for i := 0; i < b.N; i++ {
// 		bmData := BenchmarkStruct{
// 			Name:      "BenchmarkInsert",
// 			Iteration: fmt.Sprint(i),
// 		}
// 		db.Table("benchmarks").Insert(bmData).Run()
// 	}
// }

func BenchmarkSelect(b *testing.B) {
	db, _ = ydb.Connect("127.0.0.1:8008")
	for i := 0; i < b.N; i++ {
		db.Table("benchmarks").Select(&tmp).Where([][]string{
			{"iteration", "=", fmt.Sprint(i)},
		}).Run()
	}
}

// func BenchmarkUpdate(b *testing.B) {
// 	db, _ = ydb.Connect("127.0.0.1:8008")
// 	for i := 0; i < b.N; i++ {
// 		db.Table("benchmarks").Update([][]string{
// 			{"iteration", "=", fmt.Sprint(i)},
// 		},
// 			[][]string{
// 				{"iteration", fmt.Sprint(i)},
// 			},
// 		).Run()
// 	}
// }

// func BenchmarkDelete(b *testing.B) {
// 	db, _ = ydb.Connect("127.0.0.1:8008")
// 	for i := 0; i < b.N; i++ {
// 		db.Table("benchmarks").Delete([][]string{
// 			{"iteration", "=", fmt.Sprint(i)},
// 		}).Run()
// 	}
// }
