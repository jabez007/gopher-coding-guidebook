# Database Integration

Persist your in-memory todo API to a database by integrating a relational DB like SQLite or Postgres. This project builds on your previous in-memory CRUD app and adds data persistence and basic testing.

---

## 🧠 Learning Goals

- Connect a Go API to a relational database
- Migrate from in-memory data structures to persistent storage
- Perform CRUD operations using a DB layer
- Handle database errors gracefully
- Organize your code for separation of concerns
- Write tests that interact with your database

---

## 📋 Prerequisites

- Go installed (1.20+)
- Docker (for running DB containers)
- Familiarity with SQL basics
- Completed 01-crud-api (in-memory todo API)

---

## 🛠️ Task List

- [ ] Create a Task struct and matching database schema
- [ ] Set up a local database using SQLite or PostgreSQL
- [ ] Implement DB access methods (e.g., using sqlx, gorm, or database/sql)
- [ ] Replace in-memory storage with DB interactions
- [ ] Ensure all CRUD routes still function as expected
- [ ] Write tests for DB-backed handler logic

---

## 💡 Stretch Goals

- [ ] Use a database migration tool like golang-migrate
- [ ] Add created_at / updated_at timestamps
- [ ] Implement basic filtering (e.g., completed vs. active tasks)
- [ ] Add health check and DB connectivity test endpoint

---

## 📦 Stack Suggestions

- DB: SQLite (simple) or PostgreSQL (more real-world)
- DB Client: gorm, sqlx, or raw database/sql
- Router: gin, chi, or net/http
- Testing: testing, httptest
- Tooling: Docker, Docker Compose

---

## 🧪 Example Test Ideas

- Should create a new task and persist it
- Should return all tasks from the database
- Should update an existing task
- Should delete a task by ID
- Should handle DB errors gracefully

