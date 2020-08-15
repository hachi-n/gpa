package main

import (
	"github.com/hachi-n/pag/cli/pag/internal/options"
	pag "github.com/hachi-n/pag/lib/pag/show"
	"github.com/urfave/cli/v2"
)

func showCommand() *cli.Command {
	var projectType string
	return &cli.Command{
		Name:  "show",
		Usage: "show pag.database Tree",
		Flags: []cli.Flag{
			options.ProjectTypeFlag(&projectType, true, ""),
		},
		Action: func(c *cli.Context) error {
			return pag.Apply(projectType)
		},
	}
}
