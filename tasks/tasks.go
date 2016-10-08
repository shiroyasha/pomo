package tasks

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func Status(c *cli.Context) error {
	fmt.Println("status")
	return nil
}

func Add(c *cli.Context) error {
	var task = allArgs(c)

	if task == "" {
		return cli.NewExitError("Task can't be empty.", 10)
	}

	fmt.Println(task)

	return nil
}

func List(c *cli.Context) error {
	fmt.Println("list")

	return nil
}

func Remove(c *cli.Context) error {
	fmt.Println("remove")

	return nil
}

func allArgs(c *cli.Context) string {
	var args = strings.Join(c.Args(), " ")

	return strings.Trim(args, " ")
}
