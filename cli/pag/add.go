package main

import (
	"github.com/hachi-n/pag/cli/pag/internal/options"
	pag "github.com/hachi-n/pag/lib/pag/add"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	var config string
	return &cli.Command{
		Name:  "add",
		Usage: "add configFile => pag.database",
		Flags: []cli.Flag{
			options.ConfigFlag(&config, true, ""),
		},
		Action: func(c *cli.Context) error {
			return pag.Apply(config)
		},
	}
}
