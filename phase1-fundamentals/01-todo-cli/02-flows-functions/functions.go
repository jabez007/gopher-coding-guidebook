package main

import (
	"fmt"
	"math/rand"
)

// Function with parameters and return value
func addTask(title string, priority int) (string, int) {
	return fmt.Sprintf("Added task: %s (priority: %d)", title, priority), rand.Intn(10)
}

// Function with multiple return values
func completeTask(id int) (bool, string) {
	if id <= 0 {
		return false, "Invalid task ID"
	}
	return true, fmt.Sprintf("Task %d completed", id)
}

func main() {
	// Call function with arguments
	result, id := addTask("Learn Go", 1)
	fmt.Println(result)

	// Handle multiple return values
	success, message := completeTask(id)
	if success {
		fmt.Println(message)
	} else {
		fmt.Println("Error:", message)
	}

	// Ignore one return value with _
	success, _ = completeTask(0)
	fmt.Println("Task completed:", success)
}
