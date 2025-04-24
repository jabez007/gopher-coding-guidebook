package main

import (
	"fmt"
	"math/rand"
)

// addTask creates a new task with the given title and priority.
// It returns a confirmation message and a randomly generated task ID.
func addTask(title string, priority int) (string, int) {
	return fmt.Sprintf("Added task: %s (priority: %d)", title, priority), rand.Intn(10)
}

// completeTask marks the task with the given ID as completed.
// It returns success (true if the operation was successful) and
// a message (either a success message or an error description).
func completeTask(id int) (success bool, message string) {
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
