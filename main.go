package main

import (
	"fmt"
	"os"

	"github.com/kaibling/iggy/bootstrap/app"
	"github.com/urfave/cli/v2"
)

var (
	version   string
	buildTime string //nolint: gochecknoglobals
)

func main() {
	app := &cli.App{ //nolint:exhaustruct
		Name:    "iggy",
		Version: version + "-" + buildTime,
		Usage:   "application for executing generative workflows",
		Flags: []cli.Flag{
			&cli.BoolFlag{ //nolint:exhaustruct
				Name:     "api",
				Aliases:  []string{"a"},
				Usage:    "start of the web api",
				Required: false,
				Value:    true,
			},
			&cli.BoolFlag{ //nolint:exhaustruct
				Name:     "worker",
				Aliases:  []string{"w"},
				Usage:    "start of the worker",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			if !c.Bool("worker") && !c.Bool("api") {
				cli.ShowAppHelpAndExit(c, 0)
			}

			return app.Run(c.Bool("worker"), c.Bool("api"), version, buildTime)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err) //nolint: forbidigo
	}
}
