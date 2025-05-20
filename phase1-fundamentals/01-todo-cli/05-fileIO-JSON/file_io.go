package main

import (
	"fmt"
	"os"
)

func main() {
	// Write to file
	content := "This is my task list:\n1. Learn Go\n2. Build CLI app"
	err := os.WriteFile("tasks.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")

	// Read from file
	data, err := os.ReadFile("tasks.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(data))

	// Append to file
	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\n3. Write tests")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Content appended successfully")

	// Read updated file
	data, err = os.ReadFile("tasks.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Updated content:")
	fmt.Println(string(data))
}
