package main

import (
	"fmt"
)

func main() {
	// Arrays have fixed size
	var priorities [3]string
	priorities[0] = "Low"
	priorities[1] = "Medium"
	priorities[2] = "High"

	fmt.Println("Priority levels:", priorities)
	fmt.Println("Number of priorities:", len(priorities))

	// Array initialization
	categories := [4]string{"Work", "Personal", "Study", "Health"}
	fmt.Println("Categories:", categories)

	// Slices are dynamic
	tasks := []string{"Learn Go", "Build CLI app"}
	fmt.Println("Initial tasks:", tasks)

	// Append to slice
	tasks = append(tasks, "Write tests")
	fmt.Println("Updated tasks:", tasks)

	// Slice operations
	firstTwoTasks := tasks[:2]
	fmt.Println("First two tasks:", firstTwoTasks)

	// Create slice with capacity
	todoList := make([]string, 0, 10)
	fmt.Printf("Length: %d, Capacity: %d\n", len(todoList), cap(todoList))

	todoList = append(todoList, "New task")
	fmt.Printf("Length: %d, Capacity: %d\n", len(todoList), cap(todoList))
}
