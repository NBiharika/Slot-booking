package cli

import (
	"github.com/urfave/cli/v2"
)

func SetupCLI() *cli.App {
	app := &cli.App{
		Commands: cli.Commands{
			&StartCommand,
			&ScriptCommand,
		},
	}

	return app
}
