package inspector

import (
	"fmt"

	"github.com/Yadiiiig/ydb/internals/reader"
)

func database(c []string) *reader.Drivers {
	var err error
	fmt.Println(c)
	if len(c) == 1 {
		fmt.Println("Next time, try to specify a path, otherwise I can't really help you. :(")
		return nil
	}

	d, err := reader.ReadData(c[1])
	if err != nil {
		fmt.Println("Failed to read folder.")
		return nil
	}

	return d
}
