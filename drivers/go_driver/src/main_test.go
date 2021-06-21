package main

import (
	"testing"

	ydb "yadiiig.dev/ydb/go_driver/src/lib"
)

var db *ydb.Connection

func init() {
	db = ydb.Connect("127.0.0.1:8008")
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := User{
			Firstname: "Foo",
			Lastname:  "Bar",
			Email:     "foo@bar.com",
			Company:   "dev/null",
		}
		db.Table("users").Insert(user).Run()
	}
}
