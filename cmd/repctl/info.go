package main

import (
	"fmt"

	"github.com/fuzzy/repetier"
	"github.com/urfave/cli/v2"
)

func serverInfo(c *cli.Context) error {
	if len(c.String("apikey")) > 0 {
		api := repetier.NewRestClient("http", c.String("server"), c.Int("port"), c.String("apikey"))
		data := api.Info()
		if len(data.ServerName) > 0 {
			fmt.Printf("%s: %s\n", green("Server"), data.ServerName)
			fmt.Printf("%s: %s %s\n", green("Version"), data.Name, data.Version)
			fmt.Printf("%s: %d\n", green("# of Printers"), len(data.Printers))
		}
	}
	return nil
}
