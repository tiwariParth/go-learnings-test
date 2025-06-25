# Go Alert System

A Golang-based system for collecting, storing, and routing alerts with webhook integrations.

## Features

- Submit alerts via HTTP API
- Store alerts in BoltDB (embedded)
- Route alerts based on rules
- Trigger webhook actions
- Dashboard for monitoring alerts
- Containerized deployment

## Architecture

```
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│   REST API    │─┐   │  Alert Router │     │   Webhook     │
│  - Collect    │ │   │  - Filter     │     │   - HTTP      │
│  - Query      │ └──►│  - Route      │────►│   - Templates │
│  - Manage     │     │  - Prioritize │     │   - Retry     │
└───────────────┘     └───────────────┘     └───────────────┘
        │                    │                     ▲
        │                    ▼                     │
        │            ┌───────────────┐             │
        └───────────►│   Storage     │─────────────┘
                     │  - BoltDB     │
                     │  - Query      │
                     │  - Index      │
                     └───────────────┘
```

## Project Structure

```
alert-system/
├── api/              # REST API handlers
├── config/           # Configuration management
├── db/               # Database operations
├── models/           # Data models
├── router/           # Alert routing engine
├── webhook/          # Webhook execution
├── web/              # Dashboard UI
├── main.go           # Application entry point
├── Dockerfile        # Docker configuration
└── docker-compose.yml # Deployment setup
```

## API Endpoints

- `POST /api/alerts` - Submit a new alert
- `GET /api/alerts` - List all alerts
- `GET /api/alerts/{id}` - Get alert details
- `PUT /api/alerts/{id}/acknowledge` - Acknowledge an alert
- `GET /api/webhooks` - List configured webhooks
- `POST /api/webhooks` - Create a webhook
- `GET /api/stats` - Get system statistics

## Alert Format

```json
{
  "title": "Service Down",
  "severity": "critical",
  "source": "production-api",
  "description": "The API service is not responding",
  "timestamp": "2023-05-15T14:30:00Z",
  "tags": ["api", "production", "outage"],
  "data": {
    "responseCode": 500,
    "errorMessage": "Connection refused"
  }
}
```

## Getting Started

1. Clone the repository
2. Configure the application with `config.yaml`
3. Run the server
4. Access the dashboard at `/dashboard`

## Deployment

```bash
# Build and run with Docker
docker build -t alert-system .
docker run -p 8080:8080 alert-system
```
