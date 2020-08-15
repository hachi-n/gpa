package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "pag",
		Usage: "Project Generator.",
		Description: "Project Architecture Generator.",
		Commands: []*cli.Command{
			addCommand(),
			listCommand(),
			applyCommand(),
			showCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
