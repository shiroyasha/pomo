package tasks

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
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

	appendTask(task)

	return nil
}

func List(c *cli.Context) error {
	for index, task := range loadTasks() {
		fmt.Printf("%d) %s\n", index, task)
	}

	return nil
}

func Remove(c *cli.Context) error {
	index, _ := strconv.Atoi(c.Args().First())

	removeTask(index)

	return nil
}

func removeTask(index int) error {
	tasks := loadTasks()

	tasks = append(tasks[:index], tasks[index+1:]...)

	saveTasks(tasks)

	return nil
}

func appendTask(task string) {
	tasks := loadTasks()

	tasks = append(tasks, task)

	saveTasks(tasks)
}

func saveTasks(tasks []string) {
	fd, _ := os.Create("/tmp/nesto")
	defer fd.Close()

	writer := bufio.NewWriter(fd)

	for _, task := range tasks {
		fmt.Fprintf(writer, "%s\n", task)
	}

	writer.Flush()
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
