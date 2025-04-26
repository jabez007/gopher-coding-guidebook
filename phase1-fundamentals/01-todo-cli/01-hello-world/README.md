# Go Installation & First Program

## Objectives

- Set up Go development environment
- Understand Go's basic syntax and structure
- Write your first Go program

## Setup Instructions

1. Install Go
   - Download from golang.org/dl
   - Follow installation instructions for your OS
   - Verify installation: `go version`
2. Set up your editor
   - VSCode with Go extension recommended
   - GoLand is a great alternative if available
   - Configure Go tools: `go install golang.org/x/tools/...@latest`
3. Create a workspace
   ```bash
   mkdir -p ~/go-task-manager/hello
   cd ~/go-task-manager/hello
   ```

---

## ✨ Exercise 1

Create a file named `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run hello.go
```

### Key pieces

#### `package main`

- In Go, **everything must belong to a package**.
- `main` is a *special* package name — it's the one that **makes a standalone executable**.
- If you were writing a library to be imported by other code, you’d use a different package name (like `package mymath` or `package taskmanager`).
- But **if you want to run your program**, it must be in `package main`.

Think of `package main` like saying *"this is a full program you can build and run"*,
versus *"this is just a piece of code meant to be reused by other programs"*.

### `func main()`

- `func main()` is **the entry point** of a Go application — the first thing Go calls when the program runs.
- It's *required* in `package main`.
- Without it, Go will compile the file (if you run `go build`), but it won’t know what to do — it'll error if you try to `go run` it.

It's like the `main()` method in C or Java — the spot where your program officially *starts*.

- You **can** define other functions outside of `main()` and call them from `main()`.
- But `main()` **must exist** in an executable Go program.
- If you were writing a library (package other than `main`), you wouldn’t have a `func main()` at all.

---

## ✨ Exercise 2

### Concepts

Go is strongly typed, but also wants to keep things **concise** when it can figure types out for you.

#### Variable Declaration

- `var` is explicit **variable declaration**.
- You have to tell Go the **type** (`string`, `float32`, etc.) unless it can infer it.

```go
var name string = "Task Manager"
var version float32 = 0.1
```

#### Short Declaration

- `:=` is **short variable declaration**.
- It lets Go *infer* the type from the value.
- Only available **inside functions** (can’t use it at the top level of a file).

```go
isReleased := false
```

#### Constants

- `const` is for **unchanging values**.
- Good habit: constants for things that are “fixed settings” of the program (like maximum limits, API URLs, etc).

```go
const maxTasks = 100
```

> In Go, you must use all declared variables unless you intentionally ignore them with `_`.

### Practice

Create a file named `variables.go`:

```go
package main

import "fmt"

func main() {
    // Declare variables
    var name string = "Task Manager"
    var version float32 = 0.1
    
    // Short declaration
    isReleased := false
    
    fmt.Printf("Application: %s\n", name)
    fmt.Printf("Version: %.1f\n", version)
    fmt.Printf("Released: %t\n", isReleased)
    
    // Constants
    const maxTasks = 100
    fmt.Printf("Maximum tasks: %d\n", maxTasks)
}
```

- `fmt.Printf` uses formatting verbs like `%s`, `%f`, `%t`, `%d`.
  - `%s` = string
  - `%f` = float
    - `fmt.Printf("Version: %f", 1.23)` → `Version: 1.230000`
    - `fmt.Printf("Version: %.1f", 1.23)` → `Version: 1.2`
    - `fmt.Printf("Version: %05.2f", 1.23)` → `Version: 01.23`
  - `%t` = boolean
  - `%d` = integer

