package main

import (
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
	db := ydb.Connect("127.0.0.1:8008")

	user := User{
		Firstname: "Foo",
		Lastname:  "Bar",
		Email:     "foo@bar.com",
		Company:   "dev/null",
	}

	db.Table("users").Insert(user).Run()
}
