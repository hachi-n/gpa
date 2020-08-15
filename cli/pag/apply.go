package main

import (
	"github.com/hachi-n/pag/cli/pag/internal/options"
	pag "github.com/hachi-n/pag/lib/pag/apply"
	"github.com/urfave/cli/v2"
)

func applyCommand() *cli.Command {
	var projectName, output, projectType string
	return &cli.Command{
		Name:  "apply",
		Usage: "apply pag.database",
		Flags: []cli.Flag{
			options.ProjectNameFlag(&projectName, true, ""),
			options.ProjectTypeFlag(&projectType, false, ""),
			options.OutputFlag(&output, false, "."),
		},
		Action: func(c *cli.Context) error {
			return pag.Apply(projectName, output, projectType)
		},
	}
}
