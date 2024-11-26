package main

import (
	"flag"
	"fmt"
	"os"
)

const taskFilePath = "tasks.json"

func main() {
	// Define CLI flags
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", -1, "Mark a task as done (provide task ID)")
	delete := flag.Int("delete", -1, "Delete a task (provide task ID)")

	flag.Parse()

	// Load tasks from file
	tasks, err := LoadTasks(taskFilePath)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	// Handle each flag
	switch {
	case *add != "":
		tasks = append(tasks, Task{
			ID:    len(tasks) + 1,
			Title: *add,
			Done:  false,
		})
		fmt.Println("Task added:", *add)

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		} else {
			fmt.Println("Your tasks:")
			for _, task := range tasks {
				task.Display()
			}
		}

	case *done >= 1:
		found := false
		for i, task := range tasks {
			if task.ID == *done {
				tasks[i].Done = true
				fmt.Println("Task marked as done:", task.Title)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Task not found.")
		}

	case *delete >= 1:
		index := -1
		for i, task := range tasks {
			if task.ID == *delete {
				index = i
				break
			}
		}
		if index != -1 {
			fmt.Println("Deleted task:", tasks[index].Title)
			tasks = append(tasks[:index], tasks[index+1:]...)
		} else {
			fmt.Println("Task not found.")
		}

	default:
		flag.Usage()
	}

	// Save tasks back to file
	if err := SaveTasks(taskFilePath, tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}
