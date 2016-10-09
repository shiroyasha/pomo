package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shiroyasha/pomo/tasks"
	"gopkg.in/urfave/cli.v1"
)

func status(c *cli.Context) error {
	fmt.Println("status")
	return nil
}

func add(c *cli.Context) error {
	task := allArgs(c)

	if task == "" {
		return cli.NewExitError("Task can't be empty.", 10)
	}

	tasks.Add(task)

	return nil
}

func list(c *cli.Context) error {
	for index, task := range tasks.Load() {
		fmt.Printf("%d) %s\n", index, task)
	}

	return nil
}

func remove(c *cli.Context) error {
	index, _ := strconv.Atoi(c.Args().First())

	tasks.Remove(index)

	return nil
}

func allArgs(c *cli.Context) string {
	var args = strings.Join(c.Args(), " ")

	return strings.Trim(args, " ")
}

func main() {
	app := cli.NewApp()
	app.Name = "pomo"
	app.Usage = "Pomodoro timer for the Command Line"
	app.Version = "0.1.0"
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
