# ydb

ydb (Yadiiiig's Database - If you have a better name please contact me) is a hobby project where I'm trying to make my own NoSQL database. It's still a work in progress, but any feedback is always welcome.

## Main Idea

This is the current idea of how the database will look and be used when everything is finished.
You'll just have to install the executeable or build it yourself. Once you've did this you'll be able to setup an example project using a specific command.
This will setup a folder in a specified path. The folder will contain a data file, structure folder including the table structure and a backup folder.
Within the structure folder you'll be able to setup all the tables you'll need (column name and datatype). (This could also be done using the locally hosted dashboard in the future)
If that is done, you can just run the executable and point it to the database folder and it will run on the designated port.
Afterwards you can import the drivers for the language you are using and use the database.

## Features

The database currently supports basic functionality:

* Basic database functionality:
  * Select
  * Insert
  * Update
  * Delete

* Go driver [drivers/go_driver](https://github.com/Yadiiiig/ydb/tree/master/drivers/go_driver)

## Setup & run a database

Compile this project:
`go build -o ydb`

Now you can add it to your path for easy access.

Create a new database folder:
`ydb --action create --path /home/user/Documents/projectname`

Run the database:
`ydb --action run --path /home/user/Documents/projectname/database`

## Drivers

### Go driver

Example:

```go
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
```

## Current Benchmarks

These are the current benchmarks, where select-, update- and delete queries can definitly be improved.
Benchmarks are ran with one table having 10.000 entries.

```go
goos: linux
goarch: amd64
pkg: github.com/Yadiiiig/ydb/go_driver/src
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkInsert-16          9132            119983 ns/op | 0.11 ms
BenchmarkSelect-16          12798            87821 ns/op | 0.08 ms
BenchmarkUpdate-16           568           2087542 ns/op | 2.08 ms
BenchmarkDelete-16           586           1991115 ns/op | 1.99 ms
PASS
ok      github.com/Yadiiiig/ydb/go_driver/src   5.321s
```

Benchmark for 1.000.000 entries in one table.

```go
goos: linux
goarch: amd64
pkg: github.com/Yadiiiig/ydb/drivers/go_driver/src
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkSelect-16         10000            300390 ns/op | 0.30 ms
PASS
ok      github.com/Yadiiiig/ydb/drivers/go_driver/src   3.020s
```

## Upcoming features

* Finding/Implementing a search algorythm
* Adding pool connections (which is currently a bug)
* Tests
* Documentation
* Allowing integers (This shouldn't take that long to implement)
* Adding logging
* Custom Decoder so .ydb files aren't actually just json file without the .json extension (same for structure)

Starting on making a few drivers:

* Python
* Java
* Lua
* C#/.NET

Creating a local dashboard to:

* View, edit, delete data manually
* Create new database
* Create new tables
* View statistic
* ...

### People who helped me a lot

* [sn0w](https://memleak.eu/sn0w)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
