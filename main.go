package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Name: "mira",
		Version: "0.1.0-alpha",
		Authors: []*cli.Author{},
		Usage: "all tools for mirakeuri development",
		Action: func(ctx *cli.Context) error {
			log.Printf("Hello %q", ctx.Args().Get(0))
            return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
