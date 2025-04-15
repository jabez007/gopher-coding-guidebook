# Phase 2: RESTful Microservices

In this phase, you’ll build your first backend microservice using Go — exposing a REST API, persisting data to a database, and running everything inside Docker.

This is where you start building for real-world scenarios.

## 📚 Learning Goals

- Build RESTful APIs using idiomatic Go
- Use a router like `gorilla/mux` or `chi`
- Connect to a PostgreSQL or SQLite database using `database/sql` or `gorm`
- Write modular, testable services
- Build and run your app in Docker
- Use environment variables and config files
- Set up logging and error handling

## 🧱 Project

| Folder              | Description                                   |
|---------------------|-----------------------------------------------|
| `01-crud-api`       | A basic task manager REST API                 |
| `02-db-integration` | Persist and manage the task list in a databse |
| `03-k8s-deployment` | Dockerize in to a deployable microservice     |

## ✅ Goals

- Learn the anatomy of a Go-based REST service
- Build routes and handlers cleanly
- Handle DB interactions with prepared statements or ORM
- Test your services using unit and integration tests
- Dockerize and run locally with `docker-compose`

## 🚀 Stretch Challenges

- Add JWT-based authentication
- Introduce request validation and structured logging
- Add metrics with Prometheus-compatible instrumentation
- Use a `.env` file for configuration

## 📖 Helpful Resources

- https://github.com/go-chi/chi
- https://gorm.io/docs/
- https://github.com/joho/godotenv
- https://gobyexample.com

