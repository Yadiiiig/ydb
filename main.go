package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Yadiiiig/ydb/internals/background"
	"github.com/Yadiiiig/ydb/internals/creator"
	"github.com/Yadiiiig/ydb/internals/handler"
	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "action",
				Usage:       `Use 'run' to run a database or use 'create' to create a new database.`,
				Value:       "",
				DefaultText: "run",
			},
			&cli.StringFlag{
				Name:        "path",
				Usage:       `Point to parent folder of your data folder. Or point to a location to create a new data folder for a new database project. (example: user/documents/project/`,
				Value:       "",
				DefaultText: "",
			},
			&cli.IntFlag{
				Name:        "port",
				Usage:       `Specify port for the database.`,
				Value:       8008,
				DefaultText: "8008",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("action") == "run" {
			d, err := reader.ReadData(c.String("path"))
			if err != nil {
				log.Println("Failed to read folder.")
				return err
			}

			lis, err := net.Listen("tcp", ":"+fmt.Sprint(c.Int("port")))
			if err != nil {
				log.Println("failed to listen:")
				return err
			}

			background.ExitHandler(d)
			background.BackgroundUpdating(d)
			log.Println("Database is ready to Go. Hosted on port:", fmt.Sprint(c.Int("port")))
			handler.NewGrpcServer(lis, d)
		} else if c.String("action") == "create" {
			err := creator.Create(c.String("path"))
			if err != nil {
				return err
			}
			log.Println("Created a new database at: " + c.String("path"))
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
