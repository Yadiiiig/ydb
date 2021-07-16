package main

import (
	"fmt"
	"log"
	"os"

	ydb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib"
)

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Company   string
}

type Benchmarks struct {
	ID        string
	Name      string
	Iteration string
}

func main() {
	db, err := ydb.Connect("127.0.0.1:8008")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	type BenchmarkStruct struct {
		ID        string
		Name      string
		Iteration string
	}

	for i := 1; i < 1000000; i++ {
		bmData := BenchmarkStruct{
			Name:      "BenchmarkInsert",
			Iteration: fmt.Sprint(i),
		}

		_, err := db.Table("benchmarks").Insert(bmData).Run()
		if err != nil {
			fmt.Println(err)
			db, err = ydb.Connect("127.0.0.1:8008")
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
	}

	// r, err := db.Table("users").Insert(user).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(r)
	// rd, rdAmount, err := db.Table("benchmarks").Delete([][]string{
	// 	{"id", "=", "2eaa1d6b-7f7e-4578-ae50-67b987d059b9"},
	// }).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(rd, rdAmount)
	// ru, amount, err := db.Table("benchmarks").Update([][]string{
	// 	{"id", "=", "2eaa1d6b-7f7e-4578-ae50-67b987d059b9"},
	// },
	// 	[][]string{
	// 		{"id", "Hello"},
	// 	},
	// ).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(ru, amount)

	// var b []Benchmarks
	// err = db.Table("benchmarks").Select(&b).Where([][]string{
	// 	{"id", "=", "Hello"},
	// }).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, v := range b {
	// 	fmt.Println(v)
	// }

}
