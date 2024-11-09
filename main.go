package main

import (
	"fmt"
	"os"

	"github.com/kaibling/iggy/bootstrap/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "iggy",
		Usage: "application for executing generative workflows",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "api",
				Aliases:  []string{"a"},
				Usage:    "start of the web api",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "worker",
				Aliases:  []string{"w"},
				Usage:    "start of the worker",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(c.Bool("worker"), c.Bool("api"))
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
