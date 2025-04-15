# CRUD API

Build a simple RESTful API in Go that manages a todo list in memory.
This project introduces core concepts of building web servers, handling HTTP requests, working with JSON, and designing a basic CRUD API.

---

## 🧠 Learning Goals

- Understand RESTful API design using Go
- Use the `net/http` standard library or a router like `chi` or `gin`
- Handle JSON request and response bodies
- Implement basic CRUD logic (Create, Read, Update, Delete)
- Structure your code for maintainability
- Prepare for persistence in the next project

---

## 📦 Prerequisites

- Go 1.20+ installed
- Basic understanding of Go syntax and slices
- Familiarity with RESTful APIs and JSON

---

## ✅ Task List

- [ ] Set up a new Go module
- [ ] Define a `Todo` struct with fields like `ID`, `Title`, `Completed`
- [ ] Store todos in an in-memory slice or map
- [ ] Create route handlers for:
  - [ ] `POST /todos` - Create a new todo
  - [ ] `GET /todos` - List all todos
  - [ ] `GET /todos/{id}` - Get a specific todo
  - [ ] `PUT /todos/{id}` - Update a todo
  - [ ] `DELETE /todos/{id}` - Delete a todo
- [ ] Use `encoding/json` to parse requests and return responses
- [ ] Add basic logging and error handling
- [ ] Write a test for at least one route

---

## 🚀 Stretch Goals

- [ ] Add input validation (e.g., empty title)
- [ ] Implement middleware for request logging
- [ ] Use `uuid` or auto-increment ID generator
- [ ] Add filtering (e.g., completed vs. incomplete todos)

---

## 🧰 Helpful Resources

- https://gobyexample.com
- https://pkg.go.dev/net/http
- https://pkg.go.dev/encoding/json
- https://github.com/go-chi/chi (if using a router)
- https://go.dev/play/

---

This project lays the groundwork for adding database persistence in the next project.

