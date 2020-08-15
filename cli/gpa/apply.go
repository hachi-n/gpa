package main

import (
	"github.com/hachi-n/gpa/cli/gpa/internal/options"
	gpa "github.com/hachi-n/gpa/lib/gpa/apply"
	"github.com/urfave/cli/v2"
)

func applyCommand() *cli.Command {
	var projectName, output, projectType string
	return &cli.Command{
		Name:  "apply",
		Usage: "apply gpa.database",
		Flags: []cli.Flag{
			options.ProjectNameFlag(&projectName, true, ""),
			options.ProjectTypeFlag(&projectType, false, ""),
			options.OutputFlag(&output, false, "."),
		},
		Action: func(c *cli.Context) error {
			return gpa.Apply(projectName, output, projectType)
		},
	}
}
