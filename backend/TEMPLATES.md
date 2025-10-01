# Go Project Templates

This document describes the template system used for generating Go projects in the go-initializer backend.

## Overview

The template system provides standardized project structures for different types of Go applications. Each project type comes with a complete set of files and directories that follow Go best practices and community standards.

## Supported Project Types

### 1. Microservice
A complete microservice template with HTTP server, health checks, and proper project structure.

**Supported Frameworks:**
- Gin (default)
- Echo
- Fiber
- Standard HTTP

**Generated Structure:**
```
project-name/
├── cmd/project-name/
│   └── main.go                 # Application entry point
├── internal/
│   ├── server/
│   │   └── server.go          # HTTP server setup
│   ├── handler/
│   │   ├── handler.go         # Request handlers
│   │   └── health.go          # Health check endpoint
│   └── config/
│       └── config.go          # Configuration management
├── configs/
│   └── config.yaml            # Configuration file
├── api/
│   └── openapi.yaml           # OpenAPI specification
├── go.mod
├── .gitignore
├── Makefile
└── README.md
```

### 2. CLI Application
A command-line application template with proper argument parsing and subcommands.

**Supported Frameworks:**
- Cobra (recommended)
- Urfave CLI
- Kingpin
- Standard flag package

**Generated Structure:**
```
project-name/
├── cmd/project-name/
│   └── main.go                 # Application entry point
├── internal/
│   └── cli/
│       ├── root.go            # Root command
│       └── version.go         # Version command
├── go.mod
├── .gitignore
├── Makefile
└── README.md
```

### 3. API Server
Similar to microservice but with additional middleware and API-focused structure.

**Supported Frameworks:**
- Gin (default)
- Echo
- Fiber
- Chi
- Standard HTTP

**Generated Structure:**
```
project-name/
├── cmd/project-name/
│   └── main.go                 # Application entry point
├── internal/
│   ├── server/
│   │   └── server.go          # HTTP server setup
│   ├── handler/
│   │   ├── handler.go         # Request handlers
│   │   └── health.go          # Health check endpoint
│   ├── middleware/
│   │   ├── cors.go            # CORS middleware
│   │   └── auth.go            # Authentication middleware
│   └── config/
│       └── config.go          # Configuration management
├── configs/
│   └── config.yaml            # Configuration file
├── api/
│   └── openapi.yaml           # OpenAPI specification
├── go.mod
├── .gitignore
├── Makefile
└── README.md
```

### 4. Simple Project
A minimal Go project template for libraries or simple applications.

**Generated Structure:**
```
project-name/
├── cmd/project-name/
│   └── main.go                 # Application entry point
├── internal/
│   └── app/
│       └── app.go             # Application logic
├── pkg/project-name/
│   └── service.go             # Public service interface
├── go.mod
├── .gitignore
├── Makefile
└── README.md
```

## Template Features

### Common Files
All project types include these standard files:

- **README.md**: Comprehensive project documentation with setup instructions
- **go.mod**: Go module file with appropriate dependencies
- **.gitignore**: Standard Go gitignore patterns
- **Makefile**: Common build, test, and development tasks

### Framework-Specific Code
Templates generate framework-specific code for:

- HTTP server setup and configuration
- Request handlers and middleware
- Routing and endpoint definitions
- Error handling and responses
- Graceful shutdown handling

### Best Practices
Templates follow Go community best practices:

- Standard project layout (cmd/, internal/, pkg/)
- Proper error handling
- Configuration management
- Logging setup
- Testing structure
- Documentation

## Usage

The template system is used automatically when generating projects through the API:

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "projectType": "microservice",
    "goVersion": "1.22.0",
    "framework": "gin",
    "moduleName": "github.com/user/my-service",
    "name": "my-service",
    "description": "My awesome microservice"
  }'
```

## Extending Templates

To add new project types or frameworks:

1. Add the new type/framework to `types.go`
2. Create template functions in `templates.go`
3. Update the `GetTemplatesForProjectType` function
4. Add tests in `templates_test.go`

### Adding a New Project Type

```go
func getNewProjectTypeTemplates(framework string, data TemplateData) []Template {
    templates := getCommonTemplates(data)
    
    // Add project-specific templates
    templates = append(templates, []Template{
        {
            Path:    "specific/file.go",
            Content: getSpecificTemplate(framework, data),
        },
    }...)
    
    return templates
}
```

### Adding Framework Support

```go
func getFrameworkSpecificTemplate(framework string, data TemplateData) string {
    switch framework {
    case "new-framework":
        return `// New framework template content`
    default:
        return `// Default template content`
    }
}
```

## Testing

Run template tests with:

```bash
go test -v ./... -run TestTemplates
```

The test suite validates:
- Template generation for all project types
- Framework-specific code generation
- File structure completeness
- Content validation