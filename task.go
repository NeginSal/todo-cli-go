package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
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

func markDone(indexStr string) error {
	index, err := strconv.Atoi(indexStr)

	if err != nil {
		return fmt.Errorf("the Number is not valid")
	}
	index = index - 1 

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("there is no task with this number")
	}
	tasks[index].Done = true
	return saveTasks(tasks)
}

func deleteTask(indexStr string) error {
	index, err := strconv.Atoi(indexStr)

	if err != nil {
		return fmt.Errorf("number is not valid")
	}
  index = index - 1 

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("there is no task with this number")
	}

	// tasks = append(tasks[:index], tasks[index+1:]...)
	tasks = slices.Delete(tasks, index, index+1)
	return saveTasks(tasks)
}
