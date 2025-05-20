# Data Structures

This lesson introduces Go’s core data structures: arrays, slices, maps, and structs.
You'll learn how to use them in real-world scenarios and build a custom `Task` struct that forms the foundation for a CLI-based task manager.

## 🎯 Objectives

By the end of this lesson, you will:
- Understand the syntax and characteristics of arrays, slices, and maps in Go
- Learn how to define and use custom structs
- Use methods (including pointer receivers) with structs
- Begin building the data model for a CLI task manager

## ✏️ Exercise 1

📁 *File*: `collections.go`

In Go, you’ll encounter both arrays (fixed size) and slices (dynamic size).
Most of the time, you’ll work with slices, which behave more like lists or arrays in JavaScript or Python.

### 🔍 Concepts Covered:

#### 🔹 Arrays: Fixed Size

```go
var priorities [3]string
priorities[0] = "Low"
priorities[1] = "Medium"
priorities[2] = "High"
fmt.Println("Priority levels:", priorities)
```

> 🧠 **Compare**: This is like declaring a fixed-length array in TypeScript:
> `let priorities: [string, string, string] = ["Low", "Medium", "High"]`

#### 🔹 Slice Initialization

```go
tasks := []string{"Learn Go", "Build CLI app"}
fmt.Println("Tasks:", tasks)
```

> 🧠 **Compare**: This is just like `const tasks = ["Learn Go", "Build CLI app"]` in JavaScript 
> or `tasks = ["Learn Go", "Build CLI app"]` in Python.

#### 🔹 Appending to a Slice

```go
tasks = append(tasks, "Write tests")
```

> 💡 Just like `tasks.push("Write tests")` in JavaScript
> or `tasks.append("Write tests")` in Python.

#### 🔹 Slicing

```go
firstTwo := tasks[:2]
fmt.Println("First two tasks:", firstTwo)
```

> ✂️  **Compare**: This uses the same slicing syntax as Python!
> In Python: `first_two = tasks[:2]`
> In Go: `firstTwo := tasks[:2]`
> JavaScript has a similar method: `tasks.slice(0, 2)`

#### 🔹 Slices with Capacity

```go
todoList := make([]string, 0, 10)
fmt.Printf("Length: %d, Capacity: %d\n", len(todoList), cap(todoList))
```

> ⚙️ This creates an empty slice with capacity for 10 items.

##### 🧠 About `make()`

Go provides the `make()` function to create and initialize **only** slices, maps, and channels — types that need setup behind the scenes.

When you use `make()` with a slice, you can specify both its **length** and **capacity**:

```go
slice := make([]string, length, capacity)
```

- **Length**: how many elements the slice currently contains (filled with zero values).
- **Capacity**: how many elements it can hold *before* needing to resize the underlying array.

```go
todoList := make([]string, 0, 10)
// Start with 0 items, but room for 10.
```

> 💡 Why use `make()`?
> It’s useful when you know ahead of time that a slice will grow.
> Preallocating space helps improve performance by reducing internal memory copying.

> 🔍 There's no direct equivalent in JavaScript or Python. 
> Those languages grow lists dynamically, while Go gives you more control and efficiency under the hood.

### ✅ Tasks:

- Create and print an array of priorities
- Create and print an array using shorthand initialization
- Create a dynamic slice of tasks
- Append items and slice the list
- Use `make()` to explore slice capacity

> 🧠 **Self-check**: What happens when you append beyond the initial capacity? How does Go handle it?

<details>
  <summary>solution</summary>

  ```go
  package main

  import "fmt"

  func main() {
      // Arrays have fixed size
      var priorities [3]string
      priorities[0] = "Low"
      priorities[1] = "Medium"
      priorities[2] = "High"
    
      fmt.Println("Priority levels:", priorities)
      fmt.Println("Number of priorities:", len(priorities))
    
      // Array initialization
      categories := [4]string{"Work", "Personal", "Study", "Health"}
      fmt.Println("Categories:", categories)
    
      // Slices are dynamic
      tasks := []string{"Learn Go", "Build CLI app"}
      fmt.Println("Initial tasks:", tasks)
    
      // Append to slice
      tasks = append(tasks, "Write tests")
      fmt.Println("Updated tasks:", tasks)
    
      // Slice operations
      firstTwoTasks := tasks[:2]
      fmt.Println("First two tasks:", firstTwoTasks)
    
      // Create slice with capacity
      todoList := make([]string, 0, 10)
      fmt.Printf("Length: %d, Capacity: %d\n", len(todoList), cap(todoList))
    
      todoList = append(todoList, "New task")
      fmt.Printf("Length: %d, Capacity: %d\n", len(todoList), cap(todoList))
  }
  ```
</details>

---

## ✏️ Exercise 2

📁 *File*: `maps.go`

Maps in Go are key-value pairs, similar to dictionaries in Python or objects in JavaScript.
They're great for fast lookups and dynamic keys.

### 🔍 Concepts Covered:

#### 🔹 Creating a Map

```go
taskStatus := map[int]string{
    1: "Not Started",
    2: "In Progress",
    3: "Completed",
}
fmt.Println("Task statuses:", taskStatus)
```

> 🧠 **Compare**:
> JavaScript: `const taskStatus = { 1: "Not Started", 2: "In Progress", 3: "Completed" }`
> Python: `task_status = {1: "Not Started", 2: "In Progress", 3: "Completed"}`

#### 🔹 Accessing Values

```go
fmt.Println("Status 2:", taskStatus[2])
```

> 🔍 **Compare**:
> Same as `taskStatus[2]` in both JS and Python.
> But Go is stricter — if the key doesn't exist, you'll get the zero value for the type.

#### 🔹 Check If Key Exists

```go
status, exists := taskStatus[4]
if exists {
    fmt.Println("Status 4:", status)
} else {
    fmt.Println("Status 4 doesn't exist")
}
```

> 💡 In Go, this is idiomatic — you get both the value and a boolean indicating presence.
> JavaScript: `if (4 in taskStatus)` or `taskStatus.hasOwnProperty(4)`
> Python: `if 4 in task_status`

#### 🔹 Add / Delete from Map

```go
taskStatus[4] = "Blocked"
delete(taskStatus, 1)
```

> ➕ In JS: `taskStatus[4] = "Blocked"`
> ➖ To delete: `delete taskStatus[1]`
> Python is the same: `task_status[4] = "Blocked"` and `del task_status[1]`

#### 🔹 Iterate Over a Map

```go
for id, status := range taskStatus {
    fmt.Printf("ID %d: %s\n", id, status)
}
```

> 🔁 **Compare**:
> JavaScript: `for (const [id, status] of Object.entries(taskStatus))`
> Python: `for id, status in task_status.items()`

### ✅ Tasks:

- Create a `map[int]string` for task statuses
- Access a key safely and handle missing keys
- Add a new status, delete one, and print them all

> 💡 **Tip**: Map keys must be comparable types (e.g., not slices or functions)

> 🧠 **Self-check**: What’s the zero value for a string in Go? What would `taskStatus[5]` return if key 5 doesn’t exist?

<details>
  <summary>solution</summary>

  ```go
  package main

  import "fmt"

  func main() {
      // Create a map
      taskStatus := map[int]string{
          1: "Not Started",
          2: "In Progress",
          3: "Completed",
      }
    
      fmt.Println("Task statuses:", taskStatus)
    
      // Access a value
      fmt.Println("Status 2:", taskStatus[2])
    
      // Check if key exists
      status, exists := taskStatus[4]
      if exists {
          fmt.Println("Status 4:", status)
      } else {
          fmt.Println("Status 4 doesn't exist")
      }
    
      // Add a new key-value pair
      taskStatus[4] = "Blocked"
      fmt.Println("Updated statuses:", taskStatus)
    
      // Delete a key
      delete(taskStatus, 1)
      fmt.Println("After deletion:", taskStatus)
    
      // Iterate over map
      fmt.Println("All statuses:")
      for id, status := range taskStatus {
          fmt.Printf("ID %d: %s\n", id, status)
      }
  }
  ```
</details>

---

## ✏️ Exercise 3

📁 *File*: `structs.go`

Structs are Go’s way to define custom types — like classes or interfaces in TypeScript or Python.
They bundle related data together and can have methods.

### 🔍 Concepts Covered:

#### 🔹 Define a Struct

```go
type Task struct {
    ID          int
    Title       string
    Description string
    DueDate     time.Time
    Completed   bool
    Priority    int // 1: High, 2: Medium, 3: Low
}
```

> 🧱 **Compare**:
> TypeScript interface:
> ```ts
> interface Task {
>   id: number;
>   title: string;
>   description?: string;
>   dueDate: Date;
>   completed: boolean;
>   priority: 1 | 2 | 3;
> }
> ```
> Python class:
> ```python
> class Task:
>     def __init__(self, id, title, description, due_date, completed, priority):
>         self.id = id
>         ...
> ```

#### 🔹 Add Methods to Structs

```go
func (t Task) IsOverdue() bool {
    return !t.Completed && time.Now().After(t.DueDate)
}
```

> 🔍 Similar to defining a method in a class in Python/TS, using `this` or `self`.

```go
func (t *Task) Complete() {
    t.Completed = true
}
```

> 🧠 Go uses *pointer receivers* when you want to **modify the struct** (like passing by reference).
> In Python, everything is reference by default.
> In JavaScript, you'd mutate properties directly: `task.completed = true.`

#### 🔹 Using the Struct

```go
task := Task{
    ID:          1,
    Title:       "Learn Go Structs",
    Description: "Understand how to use structs in Go",
    DueDate:     time.Now().Add(24 * time.Hour),
    Completed:   false,
    Priority:    1,
}
```

> 🧠 Similar to instantiating an object in Python:
> `task = Task(1, "Learn Go Structs", ...)`
> Or creating an object in JavaScript:
> `const task = { id: 1, title: "Learn Go Structs", ... }`

#### 🔹 Slice of Structs

```go
tasks := []Task{
    task,
    {
        ID:       2,
        Title:    "Build CLI app",
        DueDate:  time.Now().Add(48 * time.Hour),
        Priority: 2,
    },
}
```

> 📋 Just like a list of objects in Python or an array of objects in JavaScript/TypeScript.

#### 🔹 Iterating Over Structs

```go
for _, t := range tasks {
    fmt.Printf("- %s (Priority: %d)\n", t.Title, t.Priority)
}
```

> 🔁 Looping is just like `for...of` in JS or `for task in tasks` in Python.

> Unlike Python or JavaScript, **Go doesn’t let you directly iterate over the fields (properties) of a struct** using something like `for key in obj`.
> This is one area where Go takes a more static and explicit approach — reflecting Go’s design philosophy of simplicity and performance over flexibility.

### ✅ Tasks:

- Create a Task struct with multiple fields
- Add a method IsOverdue() using value receiver
- Add a method Complete() using pointer receiver
- Create a slice of tasks and print them

> 💡 **Tip**: Use `fmt.Printf("%+v", obj)` to see all field values of a struct.

> 🧠 **Self-check**: Why does `Complete()` use a pointer receiver, but `IsOverdue()` does not?

<details>
  <summary>solution</summary>

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  // Task struct definition
  type Task struct {
      ID          int
      Title       string
      Description string
      DueDate     time.Time
      Completed   bool
      Priority    int // 1: High, 2: Medium, 3: Low
  }

  // Method on Task struct
  func (t Task) IsOverdue() bool {
      return !t.Completed && time.Now().After(t.DueDate)
  }

  // Method that modifies the struct (pointer receiver)
  func (t *Task) Complete() {
      t.Completed = true
  }

  func main() {
      // Create a task
      task := Task{
          ID:          1,
          Title:       "Learn Go Structs",
          Description: "Understand how to use structs in Go",
          DueDate:     time.Now().Add(24 * time.Hour),
          Completed:   false,
          Priority:    1,
      }
    
      fmt.Printf("Task: %+v\n", task)
    
      // Access fields
      fmt.Println("Title:", task.Title)
    
      // Call method
      fmt.Println("Is overdue:", task.IsOverdue())
    
      // Call method with pointer receiver
      task.Complete()
      fmt.Println("Completed:", task.Completed)
    
      // Create slice of tasks
      tasks := []Task{
          task,
          {
              ID:       2,
              Title:    "Build CLI app",
              DueDate:  time.Now().Add(48 * time.Hour),
              Priority: 2,
          },
      }
    
      fmt.Println("\nAll tasks:")
      for _, t := range tasks {
          fmt.Printf("- %s (Priority: %d)\n", t.Title, t.Priority)
      }
  }
  ```
</details>

---

## 📚 Further Practice

- **Add Tags**: Extend the `Task` struct with a `Tags []string` field.
- **Map of Tasks**: Create a `map[int]Task` to simulate a task lookup.
- **Sort by Priority**: Research how to sort a slice of structs in Go using `sort.Slice`.

## 🛠️ Bonus: Prep for CLI App

This is the groundwork for a CLI-based task manager. Start thinking about:

- How you’ll persist task data
- How to load and update task statuses
- How you might categorize or filter tasks
