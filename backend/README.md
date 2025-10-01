
# go-initializer Backend

This is the backend service for **go-initializer**, a modern web-based tool to quickly scaffold Go projects with your preferred project type, Go version, and framework/dependency.

## Features

- REST API for project scaffolding
- Generates Go project boilerplate as a downloadable ZIP
- Supports multiple project types (Microservice, CLI App, API Server, Simple Project)
- Context-aware framework/dependency selection
- Comprehensive template system with best practices
- Framework-specific code generation (Gin, Echo, Fiber, Cobra, etc.)
- Standard project layout following Go community conventions
- Written in idiomatic Go using [Gin](https://github.com/gin-gonic/gin)

## Project Structure

- `main.go` – Entry point for the backend server
- `config.go` – Configuration loading and management
- `generator.go` – Project generation logic
- `templates.go` – Template system for code generation
- `templates_test.go` – Template system tests
- `handler.go` – HTTP handlers (API endpoints)
- `router.go` – API routing setup
- `server.go` – HTTP server setup
- `types.go` – Type definitions and supported configurations
- `mod_dependencies.go` – Dependency management for generated projects
- `TEMPLATES.md` – Template system documentation

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
	"moduleName": "github.com/user/project",
	"name": "my-service",
	"description": "A Go microservice"
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
