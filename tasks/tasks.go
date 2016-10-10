package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const filePath = "/tmp/tasks"

const STATE_PENDING = "pending"
const STATE_STARTED = "started"

type Task struct {
	Description string
	State       string
}

func Start(index int) error {
	tasks := Load()

	task := &tasks[index]
	task.State = STATE_STARTED

	Save(tasks)

	return nil
}

func Add(description string) error {
	tasks := Load()

	task := Task{Description: description, State: STATE_STARTED}

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

func Init() {
	_, err := os.Stat(filePath)

	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		tasks := []Task{}

		Save(tasks)
	}
}

func Load() []Task {
	Init()

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
