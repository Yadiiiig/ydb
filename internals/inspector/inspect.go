package inspector

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Yadiiiig/ydb/internals/reader"
)

func RunInspector() {
	var d *reader.Drivers
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ydb> ")
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		command = strings.TrimSuffix(command, "\n")
		c := strings.Fields(command)

		switch c[0] {
		case "help":
			help()
		case "database":
			d = database(c)
		case "test":
			fmt.Println(d)
		case "quit":
			quit()
			return
		default:
			fmt.Println("Can't help you with this one... But if you want it, you can always contribute ;) https://github.com/Yadiiiig/ydb")
		}
		fmt.Print("ydb> ")
	}

}
