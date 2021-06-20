package main

import (
	"testing"
)

var db *Connection

func init() {
	db = connect("127.0.0.1:8008")
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := User{
			Firstname: "Foo",
			Lastname:  "Bar",
			Email:     "foo@bar.com",
			Company:   "dev/null",
		}
		db.table("users").insert(user).run()
	}
}
