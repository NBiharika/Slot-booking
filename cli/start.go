package cli

import (
	"Slot_booking/servers"
	"fmt"
	"github.com/urfave/cli/v2"
	"sort"
)

func SetupCLI() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Starting server file",
				Action: func(c *cli.Context) error {
					fmt.Println("We are starting your server")
					servers.Server()
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
