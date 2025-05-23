# 🧠 Error Handling & Command Line Arguments 

## 🎯 Objectives
By the end of this lesson, you should be able to:
- Understand how Go handles errors and why its pattern is different from other languages
- Implement and check for errors using Go’s `error` interface
- Use `os.Args` and the `flag` package to build a user-friendly CLI with subcommands
- Lay the foundation for a Task Manager CLI application

---

## 🧪 Error Handling in Go

### ✅ What is an `error` in Go?

In Go, **errors are values**, not exceptions.
This means they don’t interrupt the program like a try-catch block would in Python or JavaScript. 
Instead, Go functions return an error **explicitly**, allowing you to handle it in a very controlled and predictable way.

#### 🔁 The Pattern

Go functions that might fail will often return **two values**:

```go
result, err := someFunction()
if err != nil {
    // handle error
}
```

If there's no error, `err` will be `nil`.

#### 🧨 Simple Errors

A **simple error** is a basic error created using Go’s standard library functions, typically `errors.New`.

##### ✅ Defining a Simple Error

Use `errors.New()` to define reusable sentinel errors:

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("item not found")

func getItem(id int) (string, error) {
    if id != 42 {
        return "", ErrNotFound
    }
    return "Item 42", nil
}
```

🧩 Using a Simple Error in Calling Code

```go
func main() {
    item, err := getItem(7)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            fmt.Println("Oops! That item wasn't found.")
        } else {
            fmt.Println("An unexpected error occurred:", err)
        }
        return
    }

    fmt.Println("Retrieved:", item)
}
```

> 🔍 `errors.Is(err, ErrNotFound)` checks if `err` matches the known sentinel error, even if it's been wrapped.

#### 🎁 Wrapped Errors

A **wrapped error** includes more context about the error and keeps a reference to the underlying cause, which helps with debugging and tracing.

##### ✅ Creating a Wrapped Error

Use `fmt.Errorf()` with `%w` to wrap another error, preserving its context:

```go
package main

import (
    "fmt"
    "strconv"
)

func parseID(input string) (int, error) {
    id, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("invalid input '%s': %w", input, err)
    }
    return id, nil
}
```

🧩 Handling a Wrapped Error in Calling Code

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

func main() {
    input := "abc"
    id, err := parseID(input)
    if err != nil {
        fmt.Println("Error parsing ID:", err)

        // Unwrap and inspect inner error
        var numErr *strconv.NumError
        if errors.As(err, &numErr) {
            fmt.Println("Hint: Please enter a valid number.")
        }
        return
    }

    fmt.Println("Parsed ID:", id)
}
```

> 🔍 `errors.As(err, &numErr)` allows type assertion on the wrapped error.
> This is useful when you're working with errors from standard packages like `strconv`, `os`, `io`, etc.

#### 💡 Summary: When to Use Each

| Type of Error    | Purpose                                         | How to Handle                                                                |
|------------------|-------------------------------------------------|------------------------------------------------------------------------------|
| `errors.New`     | Sentinel values for known, predictable errors   | `errors.Is(err, ErrX)`                                                       |
| `fmt.Errorf`     | Add context while preserving the original error | `errors.Is` or `errors.As` depending on what you wrapped                     |
| Standard lib err | Returned by Go packages (e.g., `strconv`, `os`, etc.) | Use `errors.As` to extract specific error type (e.g., `*os.PathError`) |

### 🎬 Activity

1.  Create a new file called `error_handling.go`.
2.  Define a package-level custom error variable called `ErrInvalidTaskID` with the message "invalid task ID"
3.  Create the following functions:
    -  `getTask`:
       -  Takes an integer ID
       -  Return the appropriate error if the input ID is less than or equal to 0
       -  If the ID is greater than 10, return a formatted error using `fmt.Errorf()` that includes the ID in the message
       -  Otherwise, Return a string in the format "Task X" (where X is the ID) and nil error
    -  `parseID`:
       -  Takes a string and attempts to parse it into an integer ID
       -  If conversion fails, return a *wrapped* error using `fmt.Errorf()` and include the original string in your error message
       -  After successful conversion, check if the result is less than or equal to 0 and return `ErrInvalidTaskID` if so
       -  Otherwise, return the parsed integer
4.  In main, test your functions under different scenarios:
    -  Call `getTask(5)` and handle the result
       - Print the task name if successful, or print the error if it fails
    -  Call getTask(-1)
       - Check for your custom error
       - Print a user-friendly message "Please provide a valid task ID" for the custom error
       - Print the generic error message for other errors
    -  Call `getTask(100)` and print any error that occurs
    -  Call `parseID` with an invalid string
       - Print the error
       - Check if the error is of type `*strconv.NumError`, and print "Hint: Please provide a numeric ID" if it is
       - Otherwise, print "Parsed ID: X"
5.  Compile and run the code:
    ```bash
    go run error_handling.go
    ```
4.  Carefully read through the output. Try changing input values like task IDs to trigger different errors.
5.  Use `fmt.Println()` to add debug output so you can see which parts of the code are executing.

<details>
  <summary>solution</summary>

  ```go
  package main

  import (
      "errors"
      "fmt"
      "os"
      "strconv"
  )

  // Custom error
  var ErrInvalidTaskID = errors.New("invalid task ID")

  // Function that returns an error
  func getTask(id int) (string, error) {
      if id <= 0 {
          return "", ErrInvalidTaskID
      }
      if id > 10 {
          return "", fmt.Errorf("task with ID %d not found", id)
      }
      return fmt.Sprintf("Task %d", id), nil
  }

  // Function with multiple return values including error
  func parseID(s string) (int, error) {
      id, err := strconv.Atoi(s)
      if err != nil {
          return 0, fmt.Errorf("could not parse ID '%s': %w", s, err)
      }
      if id <= 0 {
          return 0, ErrInvalidTaskID
      }
      return id, nil
  }

  func main() {
      // Basic error checking
      task, err := getTask(5)
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println(task)
      }
  
      // Error from negative ID
      task, err = getTask(-1)
      if err != nil {
          if errors.Is(err, ErrInvalidTaskID) {
              fmt.Println("Please provide a positive task ID")
          } else {
              fmt.Println("Error:", err)
          }
      }
  
      // Error from large ID
      task, err = getTask(100)
      if err != nil {
          fmt.Println("Error:", err)
      }
  
      // Error wrapping
      id, err := parseID("abc")
      if err != nil {
          fmt.Println("Error:", err)
          // Check wrapped error
          var numError *strconv.NumError
          if errors.As(err, &numError) {
              fmt.Println("Hint: Please provide a numeric ID")
          }
      } else {
          fmt.Println("Parsed ID:", id)
      } 
  }
  ```
</details>

#### 🧩 Mini-Challenge

Add the following function:

```go
func markTaskComplete(id int) error
```

-   Return `ErrInvalidTaskID` if the ID is zero or negative.
-   Return `fmt.Errorf("no task with ID %d", id)` if the ID is too large.    
-   Call it from `main()` with test values like `-1`, `0`, `100`, and `5`.

✅ Bonus: Use `errors.Is` in `main()` to check for `ErrInvalidTaskID`.

### 🔄 Reflect & Review

Answer these questions in your notes or discuss with a peer:
-   How does Go’s error handling differ from `try/catch` in Python or JavaScript?
-   What are the advantages of making errors part of a function’s return signature?
-   Why might you want to “wrap” an error with more context?

---

## 🧪 Parsing Command-Line Arguments

### 🧾 Understanding `os.Args`

Go gives you direct access to command-line arguments via the `os.Args` slice.

```go
import "os"

func main() {
    fmt.Println(os.Args) // ["program", "arg1", "arg2", ...]
}
```

⚠️ Note: `os.Args[0]` is the name of your program.

### 🔧 Parsing Flags with the `flag` Package

The `flag` package in Go is part of the standard library and is used to parse **command-line options** (also called flags).
It allows you to define flags like `-title`, `-priority`, or `-due` and bind them to variables in your code.

#### 🧱 Basic Usage Pattern

```go
import "flag"

title := flag.String("title", "", "The task title")
flag.Parse()

fmt.Println("Title:", *title)
```

This line:

```go
title := flag.String("title", "", "The task title")
```

does the following:

| Part               | Description                                                                    | 
|--------------------|--------------------------------------------------------------------------------|
| `"title"`          | The name of the flag, used in the command-line like `-title="Buy milk"`        |
| `""`               | The default value if the user doesn't provide the flag                         |
| `"The task title"` | A help message shown if the user runs the program incorrectly or asks for help |

⚠️ **Important:** The variable `title` is a pointer to a string, not a string itself. 
So when you want to access the actual value, you use `*title`.

#### 📎 Example in Full Context

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    title := flag.String("title", "", "The task title")
    flag.Parse()

    if *title == "" {
        fmt.Println("Error: title is required")
        flag.PrintDefaults()
        return
    }

    fmt.Println("Adding task with title:", *title)
}
```

##### 🧪 How to Run This

If you saved that file as `add_task.go`, you can run it from the terminal like this:

```bash
go run add_task.go -title="Buy groceries"
```

Output:

```bash
Adding task with title: Buy groceries
```

If you forget the `-title` flag:

```bash
go run add_task.go
```

Output:

```bash
Error: title is required
  -title string
        The task title
```

#### ✅ Flag Types

The `flag` package supports different types of flags:

| Type       | Example Declaration                                | Command-Line Example            | 
|------------|----------------------------------------------------|---------------------------------|
| `String`   | `flag.String("title", "", "task title")`           | `-title="Buy milk"`             |
| `Int`      | `flag.Int("priority", 1, "task priority")`         | `-priority=2`                   |
| `Bool`     | `flag.Bool("completed", false, "...")`             | `-completed` (true if included) |
| `Duration` | `flag.Duration("timeout", time.Second, "timeout")` | `-timeout=5s`                   |

> 🧩 Tip: You **must** call `flag.Parse()` after defining your flags. 
> If you don’t, Go won’t populate the values based on the command-line input.

#### 📦 Subcommands with `flag.NewFlagSet`

When building CLIs like `git` or `docker`, you often use **subcommands** like:
-   `task add`
-   `task list`
-   `task delete`

The standard `flag` package parses flags from `os.Args`, but it's flat:
it doesn't understand commands like `add`, `list`, or `delete`.
So, to build subcommands with their own flags, you need a new `*flag.FlagSet` for each subcommand.

##### 🧰 Basic Structure

Here’s how flag.NewFlagSet is typically used:

```go
import (
    "flag"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("expected 'add' or 'list' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "add":
        addCmd := flag.NewFlagSet("add", flag.ExitOnError)
        title := addCmd.String("title", "", "Title of the task")
        due := addCmd.String("due", "", "Due date (YYYY-MM-DD)")

        // Parse the subcommand's flags (everything after the subcommand itself)
        addCmd.Parse(os.Args[2:])

        fmt.Println("Add command executed")
        fmt.Println("Title:", *title)
        fmt.Println("Due:", *due)

    case "list":
        listCmd := flag.NewFlagSet("list", flag.ExitOnError)
        all := listCmd.Bool("all", false, "Show all tasks")
        completed := listCmd.Bool("completed", false, "Only show completed tasks")

        listCmd.Parse(os.Args[2:])
        fmt.Println("List command executed")
        fmt.Println("All:", *all)
        fmt.Println("Completed:", *completed)

    default:
        fmt.Println("Unknown command:", os.Args[1])
        os.Exit(1)
    }
}
```

###### 🤔 Where is flag.Parse()?

Because `flag.Parse()` is only for the **default** flag set — the one used at the top level of your app.
But when you use `flag.NewFlagSet`, you're **creating an entirely new set of flags** that’s scoped to a subcommand.

So:
-  `flag.Parse()` → parses the default/global flags
-  `addCmd.Parse(...)` → parses flags specifically for the subcommand

###### 🤔 Why `Parse(os.Args[2:])`?

-  `os.Args[0]` is the program name
-  `os.Args[1]` is the subcommand
-  `os.Args[2:]` is everything after that: the flags for the subcommand

##### 🧪 How to Use It from the CLI

When you run the compiled program, it looks like this:

```bash
# Add a task
go run main.go add -title="Buy milk" -due="2025-06-01"

# List tasks
go run main.go list -all=true
```


To handle these, use `flag.NewFlagSet`:

```go
addCmd := flag.NewFlagSet("add", flag.ExitOnError)
title := addCmd.String("title", "", "Task title")

if os.Args[1] == "add" {
    addCmd.Parse(os.Args[2:])
}
```

### 🎬 Activity

1.  Create a new file called `cli_args.go`.
2.  Print all command-line arguments
3.  Check if at least one command is provided (exit with error code 1 if not)
4.  Extract and display the first command from the arguments
5.  Implement the following subcommands and related flags:
    -  `add`
       -  `title`: Task title
       -  `desc`: Task description 
       -  `due`: Due date formatted as YYYY-MM-DD
       -  `priority`: Priority level with 1 being the highest
    -  `list`
       -  `all`: To show all tasks
       -  `completed`: To show only completed tasks
       -  `today`: To show only tasks that are due today
    -  `complete`
       -  Require a task ID as a positional argument after the command
    -  `help`
6.  The implemented subcommands should have the following functionality:
    -  `add`
       -  Validate that `title` is provided
       -  Validate due date format if provided
       -  If no due date provided, default to 7 days from now
       -  Display the parsed task information
    -  `list`
       -  Display which filter(s) are active
    -  `complete`
       -  Show error if no ID provided, with usage information
       -  Display confirmation of which task is being completed
    -  `help`
       -  Display usage information for all commands
       -  Show the available options for each command
7.  Open a terminal and try the following:
    ```bash
    go run cli_args.go add -title="Finish homework" -due=2025-06-01
    go run cli_args.go list -today
    go run cli_args.go complete 3
    go run cli_args.go help
    ```
8.  If you get errors, read the output carefully. Try printing `os.Args` or using `flag.PrintDefaults()` to debug.

<details>
  <summary>solution</summary>
  ```go
  package main

  import (
      "flag"
      "fmt"
      "os"
      "time"
  )

  func main() {
      // Using os.Args directly
      fmt.Println("Arguments:", os.Args)
  
      if len(os.Args) < 2 {
          fmt.Println("No command provided")
          os.Exit(1)
      }

      // Basic subcommand handling
      command := os.Args[1]
      fmt.Println("Command:", command)

      // Using the flag package for options
      // Create flag sets for different commands
      addCmd := flag.NewFlagSet("add", flag.ExitOnError)
      addTitle := addCmd.String("title", "", "Task title (required)")
      addDesc := addCmd.String("desc", "", "Task description")
      addDue := addCmd.String("due", "", "Due date (YYYY-MM-DD)")
      addPriority := addCmd.Int("priority", 2, "Priority (1=High, 2=Medium, 3=Low)")
  
      listCmd := flag.NewFlagSet("list", flag.ExitOnError)
      listAll := listCmd.Bool("all", false, "Show all tasks")
      listCompleted := listCmd.Bool("completed", false, "Show only completed tasks")
      listToday := listCmd.Bool("today", false, "Show tasks due today")
  
      // Parse based on command
      switch command {
      case "add":
          addCmd.Parse(os.Args[2:])
          if *addTitle == "" {
              fmt.Println("Error: title is required")
              addCmd.PrintDefaults()
              os.Exit(1)
          }
  
          // Validate due date if provided
          var dueDate time.Time
          var err error
          if *addDue != "" {
              dueDate, err = time.Parse("2006-01-02", *addDue)
              if err != nil {
                  fmt.Println("Error: invalid due date format. Use YYYY-MM-DD")
                  os.Exit(1)
              }
          } else {
              dueDate = time.Now().AddDate(0, 0, 7) // Default to 1 week
          }
  
          fmt.Printf("Adding task: %s\n", *addTitle)
          fmt.Printf("Description: %s\n", *addDesc)
          fmt.Printf("Due date: %s\n", dueDate.Format("2006-01-02"))
          fmt.Printf("Priority: %d\n", *addPriority)

      case "list":
          listCmd.Parse(os.Args[2:])
          fmt.Println("Listing tasks:")
          fmt.Printf("  All: %t\n", *listAll)
          fmt.Printf("  Completed: %t\n", *listCompleted)
          fmt.Printf("  Due today: %t\n", *listToday)

      case "complete":
          if len(os.Args) < 3 {
              fmt.Println("Error: task ID required")
              fmt.Println("Usage: go run cli_args.go complete <id>")
              os.Exit(1)
          }
  
          taskID := os.Args[2]
          fmt.Printf("Marking task %s as complete\n", taskID)

      case "help":
          fmt.Println("Task Manager CLI")
          fmt.Println("Commands:")
          fmt.Println("  add -title=<title> [-desc=<description>] [-due=<YYYY-MM-DD>] [-priority=<1-3>]")
          fmt.Println("  list [-all] [-completed] [-today]")
          fmt.Println("  complete <task_id>")
          fmt.Println("  delete <task_id>")
          fmt.Println("  help")

      default:
          fmt.Printf("Unknown command: %s\n", command)
          fmt.Println("Run 'go run cli_args.go help' for usage")
          os.Exit(1)
      }
  }
  ```
 
</details>

#### 🧩 Mini-Challenges

1.  **Add JSON output**:
    -   To the `list` command, add a `--json` flag.
    -   Print a mock JSON array (don’t worry about real data yet).
2.  **Validate Priority**:
    -   In the `add` command, make sure `-priority` is between 1 and 3.
    -   If it’s not, print an error and exit.
3.  **Global `--verbose` flag**:
    -   Define a `verbose` flag early in the program.
    -   If enabled, print additional output throughout the program for debugging.

### 🔄 Reflect & Review

Answer these questions:
-   What happens if you forget to call `Parse()` on a flag set?
-   Why do CLI tools split logic into subcommands?
-   What makes a CLI "user-friendly"?
    
---

## 📘 Wrap-Up & Takeaways

-   ✅ Go favors explicit error handling — no exceptions!
-   ✅ You can build robust CLIs using `flag`, `os.Args`, and good error messages
-   ✅ This structure lays the foundation for building more complex CLI applications like a task manager

## 📚 Recommended Reading

-   📄 [Effective Go - Errors](https://go.dev/doc/effective_go#errors)
-   📦 [`flag` Package Docs](https://pkg.go.dev/flag)
-   🛠️ [Go by Example - Errors](https://gobyexample.com/errors)
-   🛠️ [Go by Example - Command-Line Arguments](https://gobyexample.com/command-line-arguments)

