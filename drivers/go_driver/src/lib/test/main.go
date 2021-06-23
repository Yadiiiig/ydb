package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("without")
	Select()
	fmt.Println("with")
	Select("a", "b", "c")
}

func Select(table ...string) {
	if table == nil {
		fmt.Println("empty")
	}
	fmt.Println(reflect.TypeOf(table))
}
