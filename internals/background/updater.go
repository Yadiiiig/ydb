package background

import (
	"fmt"
	"time"

	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/Yadiiiig/ydb/internals/utils"
)

func BackgroundUpdating(d *reader.Drivers) {
	go func() {
		a := 0
		for {
			if d.Tracker > 10 {
				utils.UpdateFile(d.OpenFile, d.Database, d.Layout, d.Path)
				d.Tracker = 0
				a = 0
			} else if a >= 5 && d.Tracker != 0 {
				fmt.Println("running")
				utils.UpdateFile(d.OpenFile, d.Database, d.Layout, d.Path)
				d.Tracker = 0
				a = 0
			} else {
				a += 1
			}
			time.Sleep(60 * time.Second)
		}
	}()
}
