package inspector

import (
	"fmt"
)

func help() {
	fmt.Println("Available commands:")
	fmt.Println("\thelp: print this help message")
	fmt.Println("\tdatabase <path>: selects a database project on a specific path")
	fmt.Println("\tquit: quit inspector")
}

func quit() {
	fmt.Println("Bye, byeee! Remember to stay hydrated!")
}
