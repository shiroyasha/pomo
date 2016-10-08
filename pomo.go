package main

import (
	"os"

	"github.com/shiroyasha/pomo/tasks"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "pomo"
	app.Usage = "Pomodoro timer for the Command Line"
	app.Action = tasks.Status

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task to the list",
			Action:  tasks.Add,
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Remove a task from the list",
			Action:  tasks.Remove,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List tasks",
			Action:  tasks.List,
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "Displays status",
			Action:  tasks.Status,
		},
	}

	app.Run(os.Args)
}
