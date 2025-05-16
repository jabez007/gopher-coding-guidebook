# 📦 File I/O and JSON in Go

## 🎯 Objectives

By the end of this lesson, you will be able to:

* 📀 Perform basic file operations (read, write, append) using Go's `ioutil` and `os` packages
* 🔐 Understand file permissions like `0644`
* 🧱 Define structs with proper JSON tags
* 🔄 Marshal and unmarshal JSON using the `encoding/json` package
* 📓 Create a simple persistence layer for a task manager app

---

## 📂 Section 1: Basic File Operations in Go

### 🧠 Concept Overview

Go provides simple tools for working with files via the `io/ioutil` (deprecated in Go 1.16 but still usable) and `os` packages. Here’s what we’ll learn:

* ✍️ Writing to a file using `ioutil.WriteFile()`
* 📖 Reading from a file using `ioutil.ReadFile()`
* ➕ Appending to a file using `os.OpenFile()`

### 🧪 Exercise 1: Basic File I/O

Create a file named `file_io.go`. Add the following snippets step by step:

#### ✍️ 1. Writing to a file

```go
content := "This is my task list:\n1. Learn Go\n2. Build CLI app"
err := ioutil.WriteFile("tasks.txt", []byte(content), 0644)
```

##### 🔐 What is `0644`?

This third argument sets the file permissions:

* It's an octal number (`0` prefix).
* `0644` means:

  * Owner can **read and write**
  * Group and Others can **read only**

| 📃 Digit | 👤 Entity | ✅ Permissions      |
| -------- | --------- | ------------------ |
| 6        | Owner     | Read + Write (4+2) |
| 4        | Group     | Read               |
| 4        | Others    | Read               |

#### 📖 2. Reading from a file

```go
data, err := ioutil.ReadFile("tasks.txt")
fmt.Println(string(data))
```

#### ➕ 3. Appending to a file

```go
file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY, 0644)
defer file.Close()
_, err = file.WriteString("\n3. Write tests")
```

##### 🧵 What's `defer` doing here?

* The `defer` keyword postpones the execution of a function until the surrounding function returns.
* Here, `defer file.Close()` ensures the file is closed **even if an error happens later**.
* This is a common and safe way to handle cleanup like closing files, network connections, or releasing locks.

🔁 **Multiple defers?** Yes! You can use defer multiple times in a function.

* Each deferred function is added to a stack.
* When the function returns, Go executes them in last-in, first-out (LIFO) order.

Example:

```go
func deferDemo() {
    defer fmt.Println("third")
    defer fmt.Println("second")
    fmt.Println("first")
}
```

Output:

```bash
first
second
third
```

📌 **Pro Tip**: Use `defer` to manage resources and avoid leaks. You'll often see it with `os.Open()`, HTTP responses, and database connections.

> 🆚 **Python’s `with` Statement**
> Go's `defer` is designed to prevent resource leaks and make cleanup safe and automatic,
> a lot like the `with` statement in Python is a context manager that automatically handles setup and teardown.

---

## 📟 Section 2: JSON in Go

### 🧠 Concept Overview

Go uses the `encoding/json` package to encode (marshal) and decode (unmarshal) data between Go structs and JSON format.

This is powerful for storing structured data like tasks, configurations, etc.

### 🔧 Defining Structs for JSON

To use structs with JSON:

* Field names **must be exported** (start with uppercase)
* Use **struct tags** to define how fields appear in JSON

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description,omitempty"`
    DueDate     time.Time `json:"due_date"`
    Completed   bool      `json:"completed"`
    Priority    int       `json:"priority"`
}
```

##### 📅 Example JSON

This is what a JSON object matching the `Task` struct might look like:

```json
{
  "id": 1,
  "title": "Finish lesson plan",
  "description": "Cover JSON marshalling",
  "due_date": "2025-05-15T00:00:00Z",
  "completed": false,
  "priority": 2
}
```

#### 💼 JSON Tag Features:

* `json:"id"`: maps to the "id" key in JSON
* `omitempty`: skips the field if it's zero or empty (e.g., empty string, 0, false)
* `time.Time` is encoded using [RFC3339](https://tools.ietf.org/html/rfc3339) by default

#### 🧱 Anatomy of a JSON-ready struct

When designing a struct for JSON:

* Use clear, camelCase-friendly field names
* Add `json` tags for every field
* Consider using `omitempty` for optional fields
* Avoid unexported fields (lowercase names) unless you want to exclude them

### 📄 Marshaling JSON

```go
data, err := json.Marshal(task)
```

Or with indentation:

```go
data, err := json.MarshalIndent(task, "", "  ")
```

### 📅 Unmarshaling JSON

```go
var task Task
err := json.Unmarshal(data, &task)
```

### ✅ Best Practices:

* Use capitalized field names for JSON visibility
* Use `omitempty` for optional fields
* Use `MarshalIndent` for readable output
* Use pointers when unmarshaling large or nested structures

---

## 🧪 Exercise 2: JSON Handling

Create a file named `json_tasks.go`. Practice defining and handling JSON:

1. 🧱 Define a `Task` struct (like above)
2. 🛠️ Create a slice of tasks
3. 📄 Marshal to JSON and print it
4. 📀 Write JSON to a file named `tasks.json`
5. 📂 Read from `tasks.json`
6. 📅 Unmarshal back to a Go slice
7. 💖 Loop through and print formatted output

---

## 🧘 Wrap-Up & Reflection

Take a few moments to reflect on what you've learned:

* 🧠 **Compare & Contrast**: How does Go’s approach to file and JSON handling differ from other languages you know (like Python, Java, or JavaScript)?
* 🛠️ **Practical Use Cases**: Where might you apply these skills in real-world projects (e.g., configuration files, local data persistence, caching)?
* 🔄 **Next Steps**: If you were to build a full task manager app, what features would you implement next? Think about features like editing, deleting, filtering tasks, or syncing with a remote server.

---

## 📚 References

* [`encoding/json` docs](https://pkg.go.dev/encoding/json)
* [`os` package docs](https://pkg.go.dev/os)
* [`ioutil` package (deprecated)](https://pkg.go.dev/io/ioutil)
* [Go File Permission Reference](https://chmod-calculator.com/)

