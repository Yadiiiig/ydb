package background

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"yadiiig.dev/ydb/internals/reader"
	"yadiiig.dev/ydb/internals/utils"
)

func ExitHandler(d *reader.Drivers) {
	c := make(chan os.Signal)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	go func() {
		<-c
		fmt.Println()
		log.Println("Do not force quit right now. You're files are still updating!")
		if d.Tracker != 0 {
			utils.UpdateFile(d.OpenFile, d.Database, d.Layout, d.Path)
		}
		os.Exit(0)
	}()
}
