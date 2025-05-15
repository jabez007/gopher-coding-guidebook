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

### 🧱 Project Setup with `go mod init`

While this project is small enough to run without modules,
you may find it helpful to initialize a Go module if you want to write tests or expand the project later.

To create a Go module in your project directory, run:

```bash
go mod init calculator
```

This creates a go.mod file that tracks dependencies and lets you use packages and testing tools more effectively.

#### 🤔 Wait, why initialize a CLI app as a module?

Go modules are not just for publishing libraries — they are how **all Go projects are managed** now.

Here's why using a module is useful **even for a CLI app**:

1. **Tooling Support**  
   Go tools like `go test`, `go get`, and even IDE features work best when a `go.mod` file is present.

2. **Dependency Management**  
   If you later decide to use a third-party package (for example, a math library or CLI parsing library), Go modules track and manage those dependencies for you.

3. **Better Organization**  
   Modules give your project a clear root directory. This helps Go understand what belongs to your project and what doesn't.

4. **Testing Support**  
   Writing and running tests is easier in a module-aware project. The `go test` command will know exactly what files and packages to test.

5. **Future-Proofing**  
   Today it’s a simple calculator — tomorrow it might become a full-featured app with subcommands and plugins. Starting with a module helps you scale up.

The `go.mod` file tells Go that:
    “This is the root of a Go project.”
    “The module name is calculator.”

> ✅ Good habit: Always run go mod init <name> when starting a new Go project.

### 🧪 Writing Unit Tests in Go

You can write tests by creating a file named `calculator_test.go` in the same directory.

Here’s an example of how to test an `Add` function:

```go
// calculator_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    got := Add(10, 5)
    want := 15

    if got != want {
        t.Errorf("Add(10, 5) = %v; want %v", got, want)
    }
}
```

To run tests:

```bash
go test
```

#### 💡 Test Tips

- Test file names must end with `_test.go`.
- Test functions must start with `Test` and accept `*testing.T`.
- You can write helper functions like `Add`, `Subtract`, etc., in `calculator.go` to make testing easier.

> 🔁 Refactor: If your logic is buried in main(), refactor the actual calculation into testable functions.
