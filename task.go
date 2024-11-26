package main

import "fmt"

// Task represents a single task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Display prints the task in a readable format
func (t Task) Display() {
	status := "Pending"
	if t.Done {
		status = "Done"
	} else {
		status = "Still Ongoing"
	}
	fmt.Printf("[%d] %s - %s\n", t.ID, t.Title, status)
}
