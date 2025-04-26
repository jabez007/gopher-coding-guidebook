# CLI Calculator

Create a **Command Line Interface** (**CLI**) calculator program in **Go** that can perform basic arithmetic between two numbers.
The calculator must accept either: 
- **symbolic operators**: (`+`, `-`, `*`, `/`) 
  or
- **word commands**: (`add`, `subtract`, `multiply`, `divide`).

> 🔥 Accepting both formats is bonus points!

## ✏️ Requirements

- Your program must be a single file: `calculator.go`
- It must accept command-line arguments to perform a calculation between two numbers.
- It must support one of the following formats (your choice):
  - **Symbolic operators example**:
    ```bash
    go run calculator.go 10 + 5
    ```
  - **Word commands example**:
    ```bash
    go run calculator.go add 10 5
    ```
- It must support these four operations:
  - Addition
  - Subtraction
  - Multiplication
  - Division
- It must **print the result** to the console.
- Your program must **handle errors** properly:
  - Not enough arguments
  - Invalid numbers
  - Unknown operations
  - Division by zero

## 🖥️ Examples

| Command                         | Output |
|---------------------------------|--------|
| `go run calculator.go 10 + 5`   |	15     |
| `go run calculator.go 10 - 5`   |	5      |
| `go run calculator.go 10 * 5`	  | 50     |
| `go run calculator.go 10 / 5`	  | 2      |
| `go run calculator.go add 10 5` |	15     |
| `go run calculator.go sub 10 5` |	5      |
| `go run calculator.go mul 10 5`	| 50     |
| `go run calculator.go div 10 5`	| 2      |

## 🔥 Bonus Opportunities

- Support **both** symbolic and word-based operations.
- Allow reversed arguments (operation first), like:
  ```bash
  go run calculator.go + 10 5
  ```
- Support **decimals** (use `float64` instead of `int`).

## 🔍 Hints

- Use `os.Args` to access command-line arguments.
- Use `strconv.Atoi()` or `strconv.ParseFloat()` to convert string arguments to numbers.
- Use a `switch` statement to select the correct operation.
- Check the number of arguments before running your code.
- Print helpful error messages when something goes wrong.
