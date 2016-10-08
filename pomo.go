package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func status(c *cli.Context) error {
	fmt.Println("status")
	return nil
}

func add(c *cli.Context) error {
	fmt.Println("add")
	return nil
}

func list(c *cli.Context) error {
	fmt.Println("list")
	return nil
}

func remove(c *cli.Context) error {
	fmt.Println("remove")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "pomo"
	app.Usage = "Pomodoro timer for the Command Line"
	app.Action = status

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task to the list",
			Action:  add,
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Remove a task from the list",
			Action:  remove,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List tasks",
			Action:  list,
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "Displays status",
			Action:  status,
		},
	}

	app.Run(os.Args)
}
