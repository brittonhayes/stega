package main

import (
	"log"
	"os"

	"github.com/brittonhayes/stega/internal/commands"
	"github.com/brittonhayes/stega/pkg/plugin"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "xxx"
	date    = "xxx"
)

func main() {
	app := &cli.App{
		Name:    "stega",
		Usage:   "Embed text into images",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:  "encode",
				Usage: "Encode text into an image",
				Args:  true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "input",
						Aliases:   []string{"i"},
						Usage:     "input file (random image from lorem picsum if not provided)",
						TakesFile: true,
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output file",
						Value:   "output.png",
					},
					&cli.StringFlag{
						Name:    "message",
						Aliases: []string{"m"},
						Usage:   "message to encode",
					},
				},
				Action: func(ctx *cli.Context) error {
					if ctx.String("input") == "" {
						return plugin.LoremPicsumEncoder().Encode()(ctx)
					}

					return commands.DefaultEncoder().Encode()(ctx)
				},
			},
			{
				Name:  "decode",
				Usage: "Decode text from an image",
				Args:  true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "input",
						Aliases:   []string{"i"},
						Usage:     "input file",
						Required:  true,
						TakesFile: true,
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output file (optional)",
					},
				},
				Action: commands.Decode(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
