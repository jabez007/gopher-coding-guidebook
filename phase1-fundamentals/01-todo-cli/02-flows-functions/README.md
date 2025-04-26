# Control Flow & Functions

## 🎯 Objectives

- Learn about control structures in Go
  - `if-else`
  - `for` loops
  - `switch` statements
- Create and use functions
  - parameters
  - return values
- Understand Go's return values
  - Work with multiple return values in Go functions
  - Handle unused return values properly

## 🧠 Concepts

### 🔹 Control Structures

- **If-Else**: Make decisions based on conditions.
- **For Loop**: Repeat actions with a loop.
  - `for` is the only loop keyword in Go — `while` and `do-while` do not exist.
- **Switch Statement**: Handle multiple possible conditions more cleanly than multiple `if-else` `if-else` statements.
  - The `switch` statement automatically breaks after each case (no `break` keyword needed like in JavaScript).
  - Instead Go has the `fallthrough` keyword which allows one condition to execute multiple cases (similar to not using `break` or `return` in JavaScript).
  - Go also allows `switch` statements without a condition, which is equivalent to `switch true`. This can be used over long `if-else` chains for complex conditions.

### 🔹 Functions

- Define reusable blocks of code.
- Functions can take **parameters** and return **one or multiple values**.
- In Go, functions can return multiple values (unlike many other languages).

---

## ✏️ Exercise 1

Write a small program that uses `if-else`, `for`, and `switch` to control the flow of execution.

1. Create a file named `control.go`
2. Start by writing a main function inside the main package.
3. Inside the `main` function:
   - **If-Else Statement**:
     - Define a variable `taskCount` and assign it a number (e.g., `5`).
     - Write an `if-else` `if-else` block:
       - If `taskCount == 0`, print "No tasks".
       - If `taskCount == 1`, print "1 task".
       - Otherwise, print "{taskCount} tasks".
   - **For Loop**:
     - print "Listing tasks:".
     - use a `for` loop to print each task, e.g. "Task 1", "Task 2", up to `taskCount`.
   - **Switch Statement**:
     - Define a variable `priority` with a value like `"high"`.
     - Write a `switch` statement:
       - If `priority` is `"low"`, print "Low priority".
       - If `priority` is `"medium"`, print "Medium priority".
       - If `priority` is `"high"`, print "High priority".
       - Otherwise, print "Unknown priority".
4. Save and run your program:
   ```bash
   go run control.go
   ```

### 🔥 BONUS

Change taskCount dynamically — e.g., ask the user to input the task count using fmt.Scanln().

<details>
  <summary>solution</summary>

  ```go
  package main

  import "fmt"

  func main() {
      // If-else statement
      taskCount := 5
    
      if taskCount == 0 {
          fmt.Println("No tasks")
      } else if taskCount == 1 {
          fmt.Println("1 task")
      } else {
          fmt.Printf("%d tasks\n", taskCount)
      }
    
      // For loop
      fmt.Println("Listing tasks:")
      for i := 1; i <= taskCount; i++ {
          fmt.Printf("Task %d\n", i)
      }
    
      // Switch statement
      priority := "high"
      switch priority {
      case "low":
          fmt.Println("Low priority")
      case "medium":
          fmt.Println("Medium priority")
      case "high":
          fmt.Println("High priority")
      default:
          fmt.Println("Unknown priority")
      }
  }   
  ```
</details>

---

## ✏️ Exercise 2

Create reusable functions for working with tasks.

1. Create a file named `functions.go`:
2. Define two functions above the `main` function:
   - `addTask` **function**:
     - Accepts two parameters: a `string` (title) and an `int` (priority).
     - Returns a single `string` that says: `"Added task: {title} (priority: {priority})"`.
   - `completeTask` **function**:
     - Accepts one parameter: an `int` (id).
     - Returns **two values**:
       - A `bool`: `true` if `id` is greater than 0, otherwise `false`.
       - A `string`: either "Task {id} completed" or "Invalid task ID".
3. Inside your `main` function:
   - Call `addTask("Learn Go", 1)` and print the returned message.
   - Call `completeTask(1)`:
     - Capture both return values (`success`, `message`).
     - If `success` is `true`, print the message.
     - Otherwise, print `"Error: {message}"`.
   - Demonstrate ignoring a return value:
     - Call `completeTask(0)`.
     - Ignore the second return value using `_` (underscore).
     - Just print whether `success` is `true` or `false`.
4. Save and run your program:
   ```bash
   go run functions.go
   ```

### 🔥 BONUS

- Modify `addTask` to assign a unique ID to each task.
- Create a `deleteTask(id int)` function that returns a message indicating if the deletion was successful.

<details>
  <summary>Solution</summary>
  

  ```go
  package main

  import "fmt"

  // Function with parameters and return value
  func addTask(title string, priority int) string {
      return fmt.Sprintf("Added task: %s (priority: %d)", title, priority)
  }

  // Function with multiple return values
  func completeTask(id int) (bool, string) {
      if id <= 0 {
          return false, "Invalid task ID"
      }
      return true, fmt.Sprintf("Task %d completed", id)
  }

  func main() {
      // Call function with arguments
      result := addTask("Learn Go", 1)
      fmt.Println(result)
    
      // Handle multiple return values
      success, message := completeTask(1)
      if success {
          fmt.Println(message)
      } else {
          fmt.Println("Error:", message)
      }
    
      // Ignore one return value with _
      success, _ = completeTask(0)
      fmt.Println("Task completed:", success)
  }
  ```

---

## 💡 Notes & Tips

- Remember, in Go **you must use** all variables unless you ignore them with `_`.
- `for` is the **only loop keyword** — there is no `while` or `do-while`.
- In `switch`, *8you don't need to add** `break` after each case — it auto-breaks.
- Functions can **return multiple values** natively — no need for special objects or structures.
