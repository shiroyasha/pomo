package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const filePath = "/tmp/tasks"

type Task struct {
	Description string
	State       string
}

func Add(description string) error {
	// tasks := Load()

	tasks := []Task{}

	task := Task{Description: description, State: "pendind"}

	tasks = append(tasks, task)

	json_tasks, _ := json.Marshal(tasks)

	fmt.Println(string(json_tasks))

	return nil

	// tasks = append(tasks, task)

	// Save(tasks)

	// return nil
}

func Remove(index int) error {
	tasks := Load()

	tasks = append(tasks[:index], tasks[index+1:]...)

	Save(tasks)

	return nil
}

func Load() []string {
	file, err := os.Open(filePath)

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

func Save(tasks []string) {
	fd, _ := os.Create(filePath)
	defer fd.Close()

	writer := bufio.NewWriter(fd)

	for _, task := range tasks {
		fmt.Fprintf(writer, "%s\n", task)
	}

	writer.Flush()
}
