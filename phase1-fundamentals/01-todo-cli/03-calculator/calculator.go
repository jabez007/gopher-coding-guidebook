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

func main() {
	// fmt.Printf("%s", os.Args)

	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operation> <number>")
		fmt.Println("Operations: +, -, *, /")
		os.Exit(1)
	}

	var num1, num2 float64
	var operation string
	var err error

	// Parse first number
	num1, err = strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		// maybe we have the operation first e.g. add 2 2
		operation = os.Args[1]

		// try to find the first number elsewhere
		num1, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println("Error parsing first number:", err)
			os.Exit(1)
		}
	} else {
		// Get operation
		operation = os.Args[2]
	}

	// Parse second number
	num2, err = strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error parsing second number:", err)
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
	case "-":
		fallthrough
	case "sub":
		result = subtract(num1, num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case "*":
		fallthrough
	case "mul":
		result = multiply(num1, num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case "/":
		fallthrough
	case "div":
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
