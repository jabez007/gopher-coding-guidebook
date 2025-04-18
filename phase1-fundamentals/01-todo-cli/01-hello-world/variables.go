package main

import "fmt"

func main() {
	// Declare variables
	var name string = "Task Manager"
	var version float32 = 0.1

	// Short declaration
	isReleased := false

	fmt.Printf("Application: %s\n", name)
	fmt.Printf("Version: %f\n", version)
	fmt.Printf("Version: %.1f\n", version)
	fmt.Printf("Version: %05.2f\n", version)
	fmt.Printf("Released: %t\n", isReleased)

	// Constants
	const maxTasks = 100
	fmt.Printf("Maximum tasks: %d\n", maxTasks)
}
