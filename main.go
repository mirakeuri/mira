package main

import (
	"log"
	"os"

	"github.com/mirakeuri/mira/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "mira",
		Version: "0.1.0-alpha",
		Authors: []*cli.Author{},
		Usage:   "all tools for mirakeuri development",
		Commands: []*cli.Command{
			commands.Rename,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
