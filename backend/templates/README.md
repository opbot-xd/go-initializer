# Go Project Templates

This directory contains the organized template system for generating Go projects.

## Structure

```
templates/
├── types.go          # Template types and main entry point
├── common.go         # Common templates (README, Makefile, etc.)
├── microservice.go   # Microservice-specific templates
├── cli.go           # CLI application templates
├── apiserver.go     # API server templates
├── simple.go        # Simple project templates
└── README.md        # This file
```

## Template Types

### 1. Common Templates (`common.go`)
Templates used across all project types:
- `README.md` - Project documentation
- `go.mod` - Go module file
- `.gitignore` - Git ignore patterns
- `Makefile` - Build and development tasks

### 2. Microservice Templates (`microservice.go`)
For HTTP-based microservices:
- Main application entry point
- HTTP server with graceful shutdown
- Health check endpoints
- Request handlers
- Configuration management
- OpenAPI specifications

### 3. CLI Application Templates (`cli.go`)
For command-line applications:
- CLI framework integration (Cobra, etc.)
- Command structure
- Version handling
- Flag parsing

### 4. API Server Templates (`apiserver.go`)
Enhanced microservice with additional middleware:
- CORS middleware
- Authentication middleware
- Extended API functionality

### 5. Simple Project Templates (`simple.go`)
Basic Go project structure:
- Minimal application setup
- Package organization
- Service layer

## Usage

The template system is used through the main `GetTemplatesForProjectType` function:

```go
templates := GetTemplatesForProjectType("microservice", "gin", templateData)
```

## Adding New Templates

### Adding a New Project Type

1. Create a new file (e.g., `newtype.go`)
2. Implement the template generation function:
   ```go
   func GetNewTypeTemplates(framework string, data TemplateData) []Template {
       templates := GetCommonTemplates(data)
       // Add specific templates
       return templates
   }
   ```
3. Update `types.go` to include the new type in `GetTemplatesForProjectType`

### Adding Framework Support

1. Update the appropriate project type file
2. Add framework-specific cases in template functions:
   ```go
   switch framework {
   case "new-framework":
       return `// New framework template`
   default:
       return `// Default template`
   }
   ```

### Template Variables

Templates use the `TemplateData` struct for variable substitution:
- `ProjectName` - Project/binary name
- `ModuleName` - Go module path
- `Description` - Project description
- `Framework` - Selected framework
- `GoVersion` - Go version
- `ProjectType` - Type of project

## Best Practices

1. **Consistency**: Follow Go community standards
2. **Framework Optimization**: Generate framework-specific code
3. **Production Ready**: Include proper error handling and configuration
4. **Documentation**: Generate comprehensive README files
5. **Testing**: Include test structure and examples

## Integration

The template system integrates with:
- `mod_dependencies.go` - For dependency management
- `types.go` - For supported configurations
- `generator.go` - For ZIP file generation
- Frontend template manager - For preview and management