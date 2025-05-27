package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Task struct with JSON field tags
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	Priority    int       `json:"priority"`
}

/*
 * The code assumes the file will be created in the current working directory.
 * In a real application, you might want to use a more robust path determination,
 * possibly using the `filepath` package to handle cross-platform path issues.
 */
const tasksJsonFilePath = "tasks.json"

func saveTasksToFile(tasks []Task, filename string) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func loadTasksFromFile(filename string) ([]Task, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func main() {
	// Create tasks
	tasks := []Task{
		{
			ID:          1,
			Title:       "Learn Go JSON",
			Description: "Understand JSON marshaling and unmarshaling",
			DueDate:     time.Now().Add(24 * time.Hour),
			Completed:   false,
			Priority:    1,
		},
		{
			ID:    2,
			Title: "Build task manager",
			// Description is omitted here to demonstrate the omitempty tag behavior
			DueDate:  time.Now().Add(72 * time.Hour),
			Priority: 2,
		},
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println("JSON data:")
	fmt.Println(string(jsonData))

	// Write JSON to file
	err = saveTasksToFile(tasks, tasksJsonFilePath)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	fmt.Println("JSON written to file successfully")

	// Read and Parse JSON from file
	loadedTasks, err := loadTasksFromFile(tasksJsonFilePath)
	if err != nil {
		fmt.Println("Error loading JSON file:", err)
		return
	}

	fmt.Println("\nLoaded tasks:")
	for _, task := range loadedTasks {
		fmt.Printf("- %s (Due: %s)\n",
			task.Title,
			// Go uses a reference date (2006-01-02) for formatting
			task.DueDate.Format("2006-01-02"),
		)
	}
}
