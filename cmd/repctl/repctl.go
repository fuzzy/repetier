package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func genericHandler(c *cli.Context) error {
	fmt.Printf("%+v\n", c)
	return nil
}

func main() {
	app := &cli.App{
		Name: "repctl",
		Usage: "Control a repetier-server instance.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "server",
				Usage: "Specify the repetier-server host.",
				Value: os.Getenv("REPETIER_HOST"),
			},
			&cli.StringFlag{
				Name: "apikey",
				Usage: "Specify the repetier-server API key.",
				Value: os.Getenv("REPETIER_APIKEY"),
			},
			&cli.IntFlag{
				Name: "port",
				Usage: "Specify the repetier-server port.",
				Value: 3344,
			},
		},
		Commands: []*cli.Command{
			{
				Name: "info",
				Usage: "Show information about the target repetier-server.",
				Action: serverInfo,
			},
			{
				Name: "status",
				Usage: "Show printer and job status.",
				Action: genericHandler,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
