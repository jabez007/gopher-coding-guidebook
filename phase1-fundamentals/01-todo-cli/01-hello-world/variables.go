package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var name string = "Task Manager"
	var version float32 = 0.1

	// Zero-value initialization
	var defaultTaskCount int
	var defaultName string
	fmt.Printf("Default task count: %d\n", defaultTaskCount) // Prints 0
	fmt.Printf("Default name: %q\n", defaultName)            // Prints ""

	// Short declaration
	isReleased := false

	// %s is for string formatting
	fmt.Printf("Application: %s\n", name)

	// %f is for default float formatting
	fmt.Printf("Version: %f\n", version)

	// %.1f specifies precision to 1 decimal place
	fmt.Printf("Version: %.1f\n", version)

	// %05.2f uses padding (width 5, zero-padded) with 2 decimal precision
	fmt.Printf("Version: %05.2f\n", version)

	// %t is for boolean formatting
	fmt.Printf("Released: %t\n", isReleased)

	// Constants
	const maxTasks = 100
	fmt.Printf("Maximum tasks: %d\n", maxTasks)
}
