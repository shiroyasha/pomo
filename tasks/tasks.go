package tasks

import (
	"encoding/json"
	"io/ioutil"
)

const filePath = "/tmp/tasks"

type Task struct {
	Description string
	State       string
}

func Add(description string) error {
	tasks := Load()

	task := Task{Description: description, State: "pendind"}

	tasks = append(tasks, task)

	Save(tasks)

	return nil
}

func Remove(index int) error {
	tasks := Load()

	tasks = append(tasks[:index], tasks[index+1:]...)

	Save(tasks)

	return nil
}

func Load() []Task {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic("Could not load tasks")
	}

	result := []Task{}

	json.Unmarshal(data, &result)

	return result
}

func Save(tasks []Task) {
	data, _ := json.Marshal(tasks)

	ioutil.WriteFile(filePath, data, 0644)
}
