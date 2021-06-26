package main

import (
	"fmt"
	"log"
	"os"

	ydb "github.com/Yadiiiig/ydb/go_driver/src/lib"
)

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Company   string
}

func main() {
	db, err := ydb.Connect("127.0.0.1:8008")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	user := User{
		Firstname: "Foo",
		Lastname:  "Bar",
		Email:     "foo@bar.com",
		Company:   "dev/null",
	}

	r, err := db.Table("users").Insert(user).Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	var users []User
	err = db.Table("users").Select(&users).Where([][]string{
		{"firstname", "=", "Foo"},
	}).Run()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range users {
		fmt.Println(v)
	}

	ru, amount, err := db.Table("users").Update([][]string{
		{"firstname", "=", "Foo"},
	},
		[][]string{
			{"firstname", "Hello"},
		},
	).Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ru, amount)

	rd, rdAmount, err := db.Table("users").Delete([][]string{
		{"firstname", "=", "Hello"},
	}).Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rd, rdAmount)
}
