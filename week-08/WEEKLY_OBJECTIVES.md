# Week 8: DevOps Tooling with Go

## Weekly Objectives

- [ ] Master file and I/O operations in Go
- [ ] Learn to parse different configuration formats (YAML, JSON, TOML)
- [ ] Execute shell commands from Go applications
- [ ] Interact with Docker using the Go SDK
- [ ] Build a CLI tool for container management

## Topics Covered

- Working with files, I/O
- YAML, JSON, TOML parsing
- Shell command execution (`os/exec`)
- Docker API (Go SDK)
- Writing your own mini `kubectl`

## Mini Project Goal

Build a **Container Cleaner CLI**:
- List running and stopped Docker containers
- Filter containers by age, status, or name
- Delete unused containers safely
- Generate reports on disk space usage
- Support for batch operations

## Exercises

1. Create a config file parser that supports multiple formats
2. Build a script executor with output capture
3. Implement a basic Docker container inspector

## Assessment

- Code review focusing on proper resource handling and error management
