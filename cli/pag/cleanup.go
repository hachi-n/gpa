package main

import (
	pag "github.com/hachi-n/pag/lib/pag/cleanup"
	"github.com/urfave/cli/v2"
)

func cleanupCommand() *cli.Command {
	return &cli.Command{
		Name:  "cleanup",
		Usage: "cleanup pag.database",
		Action: func(c *cli.Context) error {
			return pag.Apply()
		},
	}
}
