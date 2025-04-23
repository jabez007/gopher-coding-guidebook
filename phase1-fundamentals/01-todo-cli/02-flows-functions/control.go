package main

import (
	"fmt"
)

func main() {
	var taskCount int

	fmt.Print("Enter number of tasks: ")
	fmt.Scanln(&taskCount)

	// If-else statement
	if taskCount == 0 {
		fmt.Println("No tasks")
	} else if taskCount == 1 {
		fmt.Println("1 task")
	} else {
		fmt.Printf("%d tasks\n", taskCount)
	}

	// For loop
	fmt.Println("Listing tasks:")
	for i := 1; i <= taskCount; i++ {
		fmt.Printf("Task %d\n", i)
	}

	// Switch statement
	priority := "high"
	switch priority {
	case "low":
		fmt.Println("Low priority")
	case "medium":
		fmt.Println("Medium priority")
	case "high":
		fmt.Println("High priority")
	default:
		fmt.Println("Unknown priority")
	}
}
