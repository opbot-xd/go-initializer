# Go Initializer Template System

This document describes the complete template system implementation for the Go Initializer project.

## Overview

The template system has been completely refactored from a single large file into a modular, organized structure with both backend and frontend components.

## Backend Structure

### Template Organization

```
backend/
├── templates/
│   ├── types.go          # Core types and main entry point
│   ├── common.go         # Common templates (README, Makefile, etc.)
│   ├── microservice.go   # Microservice templates
│   ├── cli.go           # CLI application templates
│   ├── apiserver.go     # API server templates
│   ├── simple.go        # Simple project templates
│   └── README.md        # Template documentation
├── templates.go          # Main package integration
├── template_handler.go   # API handlers for template management
├── generator.go          # Updated to use new template system
└── router.go            # Updated with new API endpoints
```

### Key Features

1. **Modular Design**: Each project type has its own file
2. **Framework Support**: Templates adapt to selected frameworks
3. **API Integration**: RESTful endpoints for template management
4. **Backward Compatibility**: Existing generator still works

### New API Endpoints

- `POST /api/preview` - Preview templates without generating ZIP
- `GET /api/templates` - Get available template types and frameworks
- `GET /api/templates/stats` - Get template statistics

## Frontend Structure

### React TypeScript Components

```
frontend/src/
├── components/
│   ├── TemplateManager.tsx      # Main template management interface
│   ├── ProjectTypeSidebar.tsx   # Project type selection
│   ├── ProjectConfiguration.tsx # Configuration form
│   ├── TemplatePreview.tsx      # Template preview and file browser
│   └── Navigation.tsx           # App navigation
├── AppRouter.tsx               # Routing setup
├── types.ts                    # Updated with template types
└── service.ts                  # API integration
```

### Features

1. **Interactive Preview**: Browse generated files before download
2. **Real-time Configuration**: See project structure update as you type
3. **Framework-Specific Features**: UI adapts to selected project type
4. **Template Statistics**: View file counts, lines of code, etc.
5. **Responsive Design**: Works on desktop and mobile

## Template Types

### 1. Microservice
- **Frameworks**: Gin, Echo, Fiber, Standard HTTP
- **Features**: HTTP server, health checks, graceful shutdown, configuration
- **Files**: ~11 files including server, handlers, config, OpenAPI spec

### 2. CLI Application
- **Frameworks**: Cobra, Urfave, Kingpin, Standard flag
- **Features**: Command structure, subcommands, version handling
- **Files**: ~7 files including CLI structure and commands

### 3. API Server
- **Frameworks**: Gin, Echo, Fiber, Chi, Standard HTTP
- **Features**: REST API, CORS middleware, authentication, OpenAPI
- **Files**: ~13 files including middleware and enhanced API features

### 4. Simple Project
- **Framework**: Golly (minimal)
- **Features**: Basic structure, package organization, service layer
- **Files**: ~7 files with minimal Go project structure

## Usage

### Backend API

```bash
# Preview templates
curl -X POST http://localhost:8181/api/preview \
  -H "Content-Type: application/json" \
  -d '{
    "projectType": "microservice",
    "framework": "gin",
    "projectName": "my-service",
    "moduleName": "github.com/user/my-service",
    "description": "My awesome service",
    "goVersion": "1.22.0"
  }'

# Get template statistics
curl "http://localhost:8181/api/templates/stats?projectType=microservice&framework=gin"
```

### Frontend

1. Navigate to `/templates` in the web interface
2. Select project type and framework
3. Configure project details
4. Preview generated templates
5. Download complete project

## Development

### Adding New Project Types

1. Create new template file in `backend/templates/`
2. Implement template generation functions
3. Update `types.go` to include new type
4. Add frontend support in `ProjectTypeSidebar.tsx`
5. Update type definitions in `frontend/src/types.ts`

### Adding Framework Support

1. Update appropriate template file with new framework cases
2. Add framework to `backend/types.go` configuration
3. Update dependency mappings in `backend/mod_dependencies.go`
4. Test template generation

### Template Variables

All templates use the `TemplateData` struct:

```go
type TemplateData struct {
    ProjectName string  // Project/binary name
    ModuleName  string  // Go module path
    Description string  // Project description
    Framework   string  // Selected framework
    GoVersion   string  // Go version
    ProjectType string  // Type of project
}
```

## Benefits

1. **Maintainability**: Organized, modular code structure
2. **Extensibility**: Easy to add new project types and frameworks
3. **User Experience**: Interactive preview and configuration
4. **Developer Experience**: Better tooling and documentation
5. **Production Ready**: Generated projects follow best practices
6. **API-First**: RESTful design enables future integrations

## Migration

The new system is backward compatible:
- Existing `/generate` endpoint still works
- Old frontend functionality preserved
- New template system used internally
- Gradual migration path available

## Future Enhancements

1. **Custom Templates**: User-defined template support
2. **Template Marketplace**: Community template sharing
3. **Advanced Configuration**: Environment-specific settings
4. **Integration Testing**: Automated template validation
5. **Template Versioning**: Version control for templates

## Testing

```bash
# Backend tests
cd backend
go test ./templates/...

# Frontend tests
cd frontend
npm test

# Integration tests
npm run test:integration
```

This template system provides a solid foundation for generating high-quality Go projects with modern tooling and best practices.