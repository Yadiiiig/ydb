package main

import (
	"fmt"
	"log"
	"os"

	ydb "yadiiig.dev/ydb/go_driver/src/lib"
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
	fmt.Println(r, err)

	rs, err := db.Table("users").Select().Where([][]string{
		{"firstname", "=", "Foo"},
	}).Run()
	fmt.Println(rs, err)
}
