package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func parseArgs(args []string) (float64, float64, string, error) {
	var num1, num2 float64
	var operation string
	var err error

	// Try first format: <number> <operation> <number>
	num1, err = strconv.ParseFloat(args[1], 64)
	if err != nil {
		// Try second format: <operation> <number> <number>
		operation = args[1]
		num1, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			return 0, 0, "", fmt.Errorf("error parsing first number: %v", err)
		}
	} else {
		operation = args[2]
	}

	// Parse second number
	num2, err = strconv.ParseFloat(args[3], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("error parsing second number: %v", err)
	}

	return num1, num2, operation, nil
}

func main() {
	// fmt.Printf("%s", os.Args)

	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operation> <number>")
		fmt.Println("Operations: +, -, *, /")
		os.Exit(1)
	}

	num1, num2, operation, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Perform calculation
	var result float64
	switch operation {
	case "+":
		fallthrough
	case "add":
		result = add(num1, num2)
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case "-", "sub":
		result = subtract(num1, num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case "*", "mul":
		result = multiply(num1, num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case "/", "div":
		result, err = divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Invalid operation. Use +, -, *, or /")
		os.Exit(1)
	}
}
