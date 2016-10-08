package tasks

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strings"
)

func Status(c *cli.Context) error {
	fmt.Println("status")
	return nil
}

func Add(c *cli.Context) error {
	task := allArgs(c)

	if task == "" {
		return cli.NewExitError("Task can't be empty.", 10)
	}

	file, err := os.OpenFile("/tmp/nesto", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	_, err = file.WriteString(task)
	_, err = file.WriteString("\n")

	if err != nil {
		fmt.Println("ERROR:", err)

		return cli.NewExitError("Failed to save the task", 11)
	}

	file.Close()

	return nil
}

func List(c *cli.Context) error {
	file, err := os.Open("/tmp/nesto")

	defer file.Close()

	if err != nil {
		return cli.NewExitError("Could not load tasks.", 12)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	index := 1
	for scanner.Scan() {
		fmt.Printf("%d) %s\n", index, scanner.Text())

		index += 1
	}

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
