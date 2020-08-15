package main

import (
	"github.com/hachi-n/gpa/cli/gpa/internal/options"
	gpa "github.com/hachi-n/gpa/lib/gpa/show"
	"github.com/urfave/cli/v2"
)

func showCommand() *cli.Command {
	var projectType string
	return &cli.Command{
		Name:  "show",
		Usage: "show gpa.database Tree",
		Flags: []cli.Flag{
			options.ProjectTypeFlag(&projectType, true, ""),
		},
		Action: func(c *cli.Context) error {
			return gpa.Apply(projectType)
		},
	}
}
