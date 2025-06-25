# DevOps Job Runner

A Go-based system for executing and managing DevOps automation jobs via a REST API.

## Features

- Submit jobs via HTTP API
- Execute shell commands securely
- Real-time job status updates
- Job history and logs
- Job scheduling and prioritization
- Authentication and authorization
- Container-based isolation (optional)

## Architecture

```
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│   REST API    │◄────┤  Job Queue    │◄────┤  Workers      │
│   - Submit    │     │  - Priority   │     │  - Execution  │
│   - Status    │────►│  - Persistence│────►│  - Logging    │
│   - History   │     │  - Scheduling │     │  - Reporting  │
└───────────────┘     └───────────────┘     └───────────────┘
        │                    ▲                     │
        ▼                    │                     ▼
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│  Auth Service │     │  Data Store   │     │  Shell Exec   │
│  - API Keys   │     │  - Job Data   │     │  - Command    │
│  - Permissions│     │  - Logs       │     │  - Output     │
└───────────────┘     └───────────────┘     └───────────────┘
```

## Project Structure

```
job-runner/
├── api/              # REST API handlers
├── auth/             # Authentication and authorization
├── config/           # Configuration management
├── executor/         # Shell command execution
├── models/           # Data models
├── queue/            # Job queue implementation
├── storage/          # Data persistence
├── worker/           # Worker pool implementation
├── main.go           # Application entry point
├── Dockerfile        # Docker configuration
└── docker-compose.yml # Multi-container setup
```

## API Endpoints

- `POST /api/jobs` - Submit a new job
- `GET /api/jobs` - List all jobs
- `GET /api/jobs/{id}` - Get job details
- `DELETE /api/jobs/{id}` - Cancel a job
- `GET /api/jobs/{id}/logs` - Stream job logs
- `GET /api/stats` - Get system statistics

## Getting Started

1. Clone the repository
2. Configure the application
3. Run the server
4. Submit a job via the API

## Deployment

The application can be deployed as:
- Standalone binary
- Docker container
- Kubernetes deployment
