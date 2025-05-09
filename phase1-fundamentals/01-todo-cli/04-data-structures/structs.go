package main

import (
	"fmt"
	"time"
)

// Task struct definition
type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
	Priority    int // 1: High, 2: Medium, 3: Low
}

// Method on Task struct
func (t Task) IsOverdue() bool {
	return !t.Completed && time.Now().After(t.DueDate)
}

// Method that modifies the struct (pointer receiver)
func (t *Task) Complete() {
	t.Completed = true
}

func main() {
	// Create a task
	task := Task{
		ID:          1,
		Title:       "Learn Go Structs",
		Description: "Understand how to use structs in Go",
		DueDate:     time.Now().Add(24 * time.Hour),
		Completed:   false,
		Priority:    1,
	}

	fmt.Printf("Task: %+v\n", task)

	// Access fields
	fmt.Println("Title:", task.Title)

	// Call method
	fmt.Println("Is overdue:", task.IsOverdue())

	// Call method with pointer receiver
	task.Complete()
	fmt.Println("Completed:", task.Completed)

	// Create slice of tasks
	tasks := []Task{
		task,
		{
			ID:       2,
			Title:    "Build CLI app",
			DueDate:  time.Now().Add(48 * time.Hour),
			Priority: 2,
		},
	}

	fmt.Println("\nAll tasks:")
	for _, t := range tasks {
		fmt.Printf("- %s (Priority: %d)\n", t.Title, t.Priority)
	}
}
