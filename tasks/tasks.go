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
	for index, task := range loadTasks() {
		fmt.Printf("%d) %s\n", index, task)
	}

	return nil
}

func Remove(c *cli.Context) error {
	return nil
}

func loadTasks() []string {
	file, err := os.Open("/tmp/nesto")

	defer file.Close()

	if err != nil {
		panic("Could not load tasks")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := []string{}

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func allArgs(c *cli.Context) string {
	var args = strings.Join(c.Args(), " ")

	return strings.Trim(args, " ")
}
