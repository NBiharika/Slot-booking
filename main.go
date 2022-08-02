package main

import (
	"Slot_booking/cli"
	"log"
	"os"
)

func main() {
	app := cli.SetupCLI()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
