package main

import (
	"fmt"
)

func main() {
	// Create a map
	taskStatus := map[int]string{
		1: "Not Started",
		2: "In Progress",
		3: "Completed",
	}

	fmt.Println("Task statuses:", taskStatus)

	// Access a value
	fmt.Println("Status 2:", taskStatus[2])

	// Check if key exists
	status, exists := taskStatus[4]
	if exists {
		fmt.Println("Status 4:", status)
	} else {
		fmt.Println("Status 4 doesn't exist")
	}

	// Add a new key-value pair
	taskStatus[4] = "Blocked"
	fmt.Println("Updated statuses:", taskStatus)

	// Delete a key
	delete(taskStatus, 1)
	fmt.Println("After deletion:", taskStatus)

	// Iterate over map
	fmt.Println("All statuses:")
	for id, status := range taskStatus {
		fmt.Printf("ID %d: %s\n", id, status)
	}
}
