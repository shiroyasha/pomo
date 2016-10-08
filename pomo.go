package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

func allArgs(c *cli.Context) string {
	var args = strings.Join(c.Args(), " ")

	return strings.Trim(args, " ")
}

func status(c *cli.Context) error {
	fmt.Println("status")
	return nil
}

func add(c *cli.Context) error {
	var task = allArgs(c)

	if task == "" {
		return cli.NewExitError("Task can't be empty.", 10)
	}

	fmt.Println(task)

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
