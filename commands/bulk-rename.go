package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Rename = &cli.Command{
	Name:    "rename",
	Aliases: []string{"re"},
	Usage:   "bulk rename files in a folder with random or sequential naming",
	Action: func(ctx *cli.Context) error {
		fmt.Printf("renaming files in %s \n", ctx.Args().First())
		return nil
	},
}
