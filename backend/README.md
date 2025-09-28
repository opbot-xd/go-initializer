
# go-initializer Backend

This is the backend service for **go-initializer**, a modern web-based tool to quickly scaffold Go projects with your preferred project type, Go version, and framework/dependency.

## Features

- REST API for project scaffolding
- Generates Go project boilerplate as a downloadable ZIP
- Supports multiple project types (Microservice, CLI App, API Server, Simple Project)
- Context-aware framework/dependency selection
- Written in idiomatic Go using [Gin](https://github.com/gin-gonic/gin)

## Project Structure

- `main.go` – Entry point for the backend server
- `config.go` – Configuration loading and management
- `generator/` – Project scaffolding logic (to be extended)
- `handler/` – HTTP handlers (API endpoints)
- `router/` – API routing setup
- `server/` – HTTP server setup

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (1.20+ recommended)

### Install Dependencies

```sh
go mod tidy
```

### Run the Backend Server

```sh
go run main.go
```

By default, the server listens on the port specified in your configuration file.

### API Endpoint

- `POST /api/generate`  
Accepts a JSON payload describing the project to generate and returns a ZIP file.

#### Example Request

```json
{
	"projectType": "microservice",
	"goVersion": "1.22.0",
	"framework": "Gin",
	"moduleName": "github.com/example/project",
	"name": "my-service",
	"description": "A sample Go microservice"
}
```

#### Example Response

- Content-Type: `application/zip`
- Content-Disposition: `attachment; filename=project.zip`

## Development

- Use `go fmt ./...` to format code.
- Follow the [contribution guidelines](../CONTRIBUTING.md).

## Configuration

Configuration is loaded from a YAML file. See [`config.go`](config.go) for details.

## License

[MIT](../LICENSE)
