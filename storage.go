package main

import (
	"encoding/json"
	"os"
)

// SaveTasks writes tasks to a JSON file
func SaveTasks(filePath string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// LoadTasks reads tasks from a JSON file
func LoadTasks(filePath string) ([]Task, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
