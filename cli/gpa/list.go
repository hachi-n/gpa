package main

import (
	gpa "github.com/hachi-n/gpa/lib/gpa/list"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list gpa.database",
		Action: func(c *cli.Context) error {
			return gpa.Apply()
		},
	}
}
