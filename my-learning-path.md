# GoLang Mastery — DevOps Edition

Welcome to your personalized GoLang study roadmap! This guide will help you go from fundamentals to production-level Go code, with a specific focus on DevOps, tooling, and infrastructure development.

---

## 📅 Duration

**Total Duration:** 12 Weeks  
**Effort:** 7 days/week  
**Editor:** VS Code  
**Mode:** Project-based with theory, exercises, and tests

---

## 🧭 Table of Contents

1. [Week 1–2: Go Fundamentals](#week-1-2-go-fundamentals)
2. [Week 3: Intermediate Concepts](#week-3-intermediate-concepts)
3. [Week 4–5: Concurrency & Channels](#week-4-5-concurrency--channels)
4. [Week 6: CLI Tooling](#week-6-cli-tooling)
5. [Week 7: Network Programming](#week-7-network-programming)
6. [Week 8: DevOps Tooling with Go](#week-8-devops-tooling-with-go)
7. [Week 9–10: Testing, Error Handling, Logging](#week-9-10-testing-error-handling-logging)
8. [Week 11: Real-World Project I](#week-11-real-world-project-i)
9. [Week 12: Real-World Project II + Certification Test](#week-12-real-world-project-ii--certification-test)

---

## Week 1–2: Go Fundamentals

**Topics:**
- Go Setup & Workspace
- `main.go` structure, `package main`
- Variables, Constants
- Data Types & Structs
- Arrays, Slices, Maps
- Control Structures (if, switch, for)

**📘 Exercises:**
- Convert a simple JavaScript app to Go
- Implement a JSON parser using structs

**🛠️ Mini Project:**  
Build a **Basic ToDo CLI Tool** storing tasks in memory

**🧪 Test:**  
20-question MCQ + 2 small code writing tasks

---

## Week 3: Intermediate Concepts

**Topics:**
- Functions and methods
- Pointers and references
- Struct composition and embedding
- Interfaces and polymorphism

**🛠️ Mini Project:**  
Build a **Go-based Configuration Loader** (like `.env` parser)

---

## Week 4–5: Concurrency & Channels

**Topics:**
- Goroutines
- Channels (buffered/unbuffered)
- Select, `sync`, `WaitGroup`, `Mutex`
- Worker pools and fan-out/fan-in patterns

**🛠️ Mini Project:**  
Build a **Parallel URL Checker Tool** with concurrent workers

**🧪 Test:**  
Concurrency patterns MCQs + live coding simulation

---

## Week 6: CLI Tooling

**Topics:**
- Flag parsing (`flag` package)
- Cobra & urfave/cli packages
- Building multi-command CLI apps
- Packaging & distributing CLI tools

**🛠️ Mini Project:**  
Build a **GitHub Repo Manager CLI** using GitHub API

---

## Week 7: Network Programming

**Topics:**
- TCP, UDP clients and servers
- HTTP client/server
- REST APIs with `net/http`
- JSON Marshalling & Unmarshalling

**🛠️ Mini Project:**  
Build a **Simple HTTP Uptime Monitoring Server**

---

## Week 8: DevOps Tooling with Go

**Topics:**
- Working with files, IO
- YAML, JSON, TOML parsing
- Shell command execution (`os/exec`)
- Docker API (Go SDK)
- Writing your own mini `kubectl`

**🛠️ Mini Project:**  
Build a **Container Cleaner CLI**: List & delete unused Docker containers

---

## Week 9–10: Testing, Error Handling, Logging

**Topics:**
- Unit testing with `testing` package
- Table-driven tests
- Test coverage and mocking
- Panic vs error
- Error wrapping and propagation
- Logging best practices (`log`, `zap`, `logrus`)

**🛠️ Mini Project:**  
Refactor older projects to include proper tests and logging

**🧪 Test:**  
Mocking, structured logging, error propagation

---

## Week 11: Real-World Project I

**🧠 Goal:** Apply everything you've learned into a larger app

**🛠️ Project:**  
Build a **DevOps Job Runner**:
- Accept jobs from HTTP API
- Execute shell commands
- Return logs and statuses
- Optional: Dockerize the app

---

## Week 12: Real-World Project II + Certification Test

**🛠️ Capstone Project:**  
Build a **Golang-based Alert System**:
- Accept alerts via REST
- Store them in BoltDB or SQLite
- Trigger webhook actions
- Deploy with Docker

**📜 Final Test:**  
- 40 MCQs + 2 coding challenges + Project Review

---

## 🧰 Tools You'll Use

- VS Code with Go extension
- Docker (for containerized projects)
- Cobra / urfave/cli (CLI)
- GoReleaser (optional)
- HTTPie or Postman for testing APIs
- GitHub for version control & practice

---

## 📚 Learning Resources

- [Go by Example](https://gobyexample.com/)
- [The Go Programming Language Tour](https://tour.golang.org/)
- [GoDoc Documentation](https://pkg.go.dev/)
- [Practical Go Patterns](https://www.practical-go-lessons.com/)

---

## 🧠 Tip: Daily Habit

- Spend 15 minutes reading Go code from GitHub
- Solve at least 1 exercise every day
- Keep a `go-notes.md` file to write down patterns, gotchas, and tips

---

## 🏁 Let's Go 🚀

Clone this repo, create a `go-learnings` folder, and start from `week-01`.

Happy Go-coding! 🐹
