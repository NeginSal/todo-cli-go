package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"stauts"`
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

func addTask(title string, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		Title:       title,
		Description: description,
		Done:        false,
	}

	tasks = append(tasks, newTask)

	err = saveTasks(tasks)
	return err
}

func listTasks() error {
	tasks, err := loadTasks()

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("There is no task!")
		return nil
	}

	for i, task := range tasks {
		status := "✗"
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
	return nil
}
