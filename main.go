package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("🎯 TODO CLI")

	//load tasks from file
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// if thers is no tasks ,show a message
	if len(tasks) == 0 {
		fmt.Println("No tasks availabel")
	} else {
		_ = listTasks()
	}

	// show available commands
	fmt.Println("\n please enter a command: add | list |done | delete |exit")

	//start interactive input mode
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()

		// split input into command and arguments
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "add":
			// if not enough arguments provided
			if len(args) < 3 {
				fmt.Println("usage: add <title>")
				continue
			}

			title := strings.Join(args[1:], " ")
			fmt.Println("📝 Enter a description for this task:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			description := scanner.Text()

			err := addTask(title, description)
			if err != nil {
				fmt.Println("error adding task:", err)
			} else {
				fmt.Println("Task added")
			}

		case "list":
			// list all tasks
			_ = listTasks()

		case "done":
			// mark a task as done
			if len(args) < 2 {
				fmt.Println("usage: done <task-number>")
				continue
			}
			err := markDone(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("task marked as done")
			}

		case "delete":
			// Delete a task
			if len(args) < 2 {
				fmt.Println("usage:delete <task-number>")
				continue
			}
			err := deleteTask(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("task deleted")
			}

		case "exit":
			//exit the program
			fmt.Println("Exiting the program 👋")

		default:
			// Invalid command entered
			fmt.Println("Invalid command! Please use: add | list | done | delete | exit")
		}

	}
}
