package main

import (
	"fmt"
)

func main() {
	var taskCount int

	fmt.Print("Enter number of tasks: ")
	_, err := fmt.Scanln(&taskCount)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

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
		fallthrough // This would execute the high priority case as well
	case "high":
		fmt.Println("High priority")
	default:
		fmt.Println("Unknown priority")
	}

	// Switch with no condition (equivalent to switch true)
	switch {
	case taskCount > 10:
		fmt.Println("That's a lot of tasks!")
	case taskCount > 5:
		fmt.Println("You're quite busy")
	case taskCount > 0:
		fmt.Println("You have a manageable workload")
	default:
		fmt.Println("Nothing to do")
	}
}
