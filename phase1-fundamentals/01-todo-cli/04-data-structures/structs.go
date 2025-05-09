package main

import (
	"fmt"
	"time"
)

// Priority levels
const (
	PriorityHigh   = 1
	PriorityMedium = 2
	PriorityLow    = 3
)

// Task struct definition
type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
	Priority    int // See priority constants
}

// Method on Task struct
func (t Task) IsOverdue() bool {
	return !t.Completed && time.Now().After(t.DueDate)
}

// Method that modifies the struct (pointer receiver)
func (t *Task) Complete() {
	if !t.Completed {
		t.Completed = true
	}
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
			ID:    2,
			Title: "Build CLI app",
			// Description is missing and will be initialized to ""
			DueDate: time.Now().Add(48 * time.Hour),
			// Completed is missing and will be initialized to false
			Priority: 2,
		},
	}

	fmt.Println("\nAll tasks:")
	for _, t := range tasks {
		fmt.Printf("- %s (Priority: %d)\n", t.Title, t.Priority)
	}
}
