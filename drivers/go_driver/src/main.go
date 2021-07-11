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

	// user := User{
	// 	Firstname: "Foo",
	// 	Lastname:  "Bar",
	// 	Email:     "foo@bar.com",
	// 	Company:   "dev/null",
	// }

	// r, err := db.Table("users").Insert(user).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(r)

	var b []Benchmarks
	err = db.Table("benchmarks").Select(&b).Where([][]string{
		{"id", "=", "e45d3e4b-8d53-4907-860b-7c48d727a5ff"},
	}).Run()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	// ru, amount, err := db.Table("users").Update([][]string{
	// 	{"firstname", "=", "Foo"},
	// },
	// 	[][]string{
	// 		{"firstname", "Hello"},
	// 	},
	// ).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(ru, amount)

	// rd, rdAmount, err := db.Table("users").Delete([][]string{
	// 	{"firstname", "=", "Hello"},
	// }).Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(rd, rdAmount)
}
