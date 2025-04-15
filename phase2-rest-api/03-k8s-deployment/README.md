# Kubernetes Deployment

Take your previous REST API and fully containerize it using Docker.
Then deploy it to a local Kubernetes cluster (like Minikube or Kind).
This project introduces you to Docker, Docker Compose, and Kubernetes concepts.

---

## 🚀 Learning Goals

- Create a production-ready Dockerfile for your Go API
- Use Docker Compose to run your app with a database
- Write Kubernetes manifests for deployment and services
- Test your app running in a local K8s cluster
- (Optional) Add Skaffold for rapid dev experience

---

## 📋 Prerequisites

- Go 1.20+
- Docker installed
- Kubernetes CLI (kubectl)
- Minikube or Kind (or Docker Desktop with K8s enabled)

---

## 🛠 Task List

- [ ] Write a multi-stage Dockerfile for the Go app
- [ ] Create a .dockerignore file
- [ ] Create a docker-compose.yml to run the app with a database
- [ ] Verify everything works locally in Docker
- [ ] Create Kubernetes manifests:
  - [ ] Deployment
  - [ ] Service
  - [ ] Secret or ConfigMap (for DB config)
  - [ ] PersistentVolume (if needed)
- [ ] Deploy the app to Minikube and test via kubectl port-forward

---

## 💡 Stretch Goals

- [ ] Add Skaffold support for live rebuilds
- [ ] Add health/liveness/readiness probes to the Deployment
- [ ] Add Prometheus metrics or structured logging
- [ ] Push Docker image to Docker Hub or GitHub Container Registry
- [ ] Deploy to a cloud-hosted K8s cluster (like GKE or EKS)

---

## 📦 Stack Suggestions

- HTTP Router: chi, gin, or net/http
- DB: SQLite or PostgreSQL
- Container tools: Docker, docker-compose, kubectl, Minikube
- Dev tools: skaffold, tilt (optional)

