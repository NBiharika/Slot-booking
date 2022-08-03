package cli

import (
	"Slot_booking/servers"
	"Slot_booking/start_up"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"time"
)

var StartCommand = cli.Command{
	Name:  "start",
	Usage: "Starting server file",
	Action: func(c *cli.Context) error {
		fmt.Println("We are starting your server")
		servers.Server()
		return nil
	},
}

var ScriptCommand = cli.Command{
	Name:  "script",
	Usage: "used to run scripts",
	Subcommands: cli.Commands{
		&addSlot,
	},
}

var addSlot = cli.Command{
	Name:  "add-slot",
	Usage: "Adding a slot",
	Action: func(c *cli.Context) error {
		fmt.Println("We are adding a slot")
		start_up.Initialize()
		m := make(map[string]interface{})
		nextTime := time.Now().Add(6 * 24 * time.Hour)
		m["date"] = nextTime.Format("2006-01-02")
		err := start_up.SlotController.AddSlot(&gin.Context{}, m)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("slot added successfully")
		}
		return nil
	},
}
