package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
	status      bool   `json:"status"`
}

func loadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return tasks, nil
	}

	err = json.Unmarshal(file, &tasks)

	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}
