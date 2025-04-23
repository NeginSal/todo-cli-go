package main

import (
	"bufio"	
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	color.Cyan("🎯 TODO CLI")

	//load tasks from file
	tasks, err := loadTasks()
	if err != nil {
		color.Red("Error loading tasks:", err)
		return
	}

	// if thers is no tasks ,show a message
	if len(tasks) == 0 {
		color.Yellow("No tasks availabel")
	} else {
		_ = listTasks()
	}

	// show available commands
	color.Blue("\n please enter a command: add | list |done | delete |exit")

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
				color.White("usage:  add <title>")
				continue
			}

			title := strings.Join(args[1:], " ")
			color.Blue("📝 Enter a description for this task:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			description := scanner.Text()

			err := addTask(title, description)
			if err != nil {
				color.Red("error adding task:", err)
			} else {
				color.Green("Task added")
			}

		case "list":
			// list all tasks
			_ = listTasks()

		case "done":
			// mark a task as done
			if len(args) < 2 {
				color.White("usage:  done <task-number>")
				continue
			}
			err := markDone(args[1])
			if err != nil {
				color.Red("error", err)
			} else {
				color.Green("task marked as done")
			}

		case "delete":
			// Delete a task
			if len(args) < 2 {
				color.White("usage:  delete <task-number>")
				continue
			}
			err := deleteTask(args[1])
			if err != nil {
				color.Red("error", err)
			} else {
				color.Green("task deleted")
			}

		case "exit":
			//exit the program
			color.White("Exiting the program 👋")
			os.Exit(0)

		default:
			// Invalid command entered
			color.Blue("Invalid command! Please use: add | list | done | delete | exit")
		}

	}
}
