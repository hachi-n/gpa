package options

import "github.com/urfave/cli/v2"

func ProjectNameFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "project",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

func OutputFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "output",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

func ConfigFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "config",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

func ProjectTypeFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "type",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}
