package main

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Company   string
}

func main() {
	db := connect("127.0.0.1:8008")

	user := User{
		Firstname: "Foo",
		Lastname:  "Bar",
		Email:     "foo@bar.com",
		Company:   "dev/null",
	}

	db.table("users").insert(user).run()
}
