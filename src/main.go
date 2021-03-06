package main

import (
	"log"
	"os"

	"github.com/bongster/golang_20210228/api"
	"github.com/bongster/golang_20210228/migration"
	"github.com/bongster/golang_20210228/util"
	"github.com/bongster/golang_20210228/worker"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "services",
		Usage: "service",
		Commands: []*cli.Command{
			{
				Name: "api",
				Action: func(c *cli.Context) error {
					api.Run()
					return nil
				},
			},
			{
				Name: "worker",
				Action: func(c *cli.Context) error {
					worker.Run()
					return nil
				},
			},
			{
				Name: "migration",
				Action: func(c *cli.Context) error {
					config, _ := util.LoadConfig(".")
					migration.Run(config.DBDriver, config.DBSource, "file://./db/migration")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
