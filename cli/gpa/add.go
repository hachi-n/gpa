package main

import (
	"github.com/hachi-n/gpa/cli/gpa/internal/options"
	gpa "github.com/hachi-n/gpa/lib/gpa/add"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	var config string
	return &cli.Command{
		Name:  "add",
		Usage: "add configFile => gpa.database",
		Flags: []cli.Flag{
			options.ConfigFlag(&config, true, ""),
		},
		Action: func(c *cli.Context) error {
			return gpa.Apply(config)
		},
	}
}
