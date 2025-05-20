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
			ID:       2,
			Title:    "Build task manager",
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
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	fmt.Println("JSON written to file successfully")

	// Read JSON from file
	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse JSON
	var loadedTasks []Task
	err = json.Unmarshal(fileData, &loadedTasks)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("\nLoaded tasks:")
	for _, task := range loadedTasks {
		fmt.Printf("- %s (Due: %s)\n",
			task.Title,
			task.DueDate.Format("2006-01-02"),
		)
	}
}
