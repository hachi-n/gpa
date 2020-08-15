package main

import (
	pag "github.com/hachi-n/pag/lib/pag/list"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list pag.database",
		Action: func(c *cli.Context) error {
			return pag.Apply()
		},
	}
}
