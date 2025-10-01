package templates

import "fmt"

// GetCommonTemplates returns templates used across all project types
func GetCommonTemplates(data TemplateData) []Template {
	return []Template{
		{
			Path:    "README.md",
			Content: GetReadmeTemplate(data),
		},
		{
			Path:    "go.mod",
			Content: GetGoModTemplate(data),
		},
		{
			Path:    ".gitignore",
			Content: GetGitignoreTemplate(),
		},
		{
			Path:    "Makefile",
			Content: GetMakefileTemplate(data),
		},
	}
}

func GetReadmeTemplate(data TemplateData) string {
	return fmt.Sprintf(`# %s

%s

## Getting Started

### Prerequisites

- Go %s or later
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
   `+"```bash"+`
   git clone <repository-url>
   cd %s
   `+"```"+`

2. Install dependencies:
   `+"```bash"+`
   go mod tidy
   `+"```"+`

### Running the Application

`+"```bash"+`
go run cmd/%s/main.go
`+"```"+`

Or using Make:
`+"```bash"+`
make run
`+"```"+`

### Building

`+"```bash"+`
go build -o bin/%s cmd/%s/main.go
`+"```"+`

Or using Make:
`+"```bash"+`
make build
`+"```"+`

### Testing

`+"```bash"+`
go test ./...
`+"```"+`

Or using Make:
`+"```bash"+`
make test
`+"```"+`

## Project Structure

`+"```"+`
.
├── cmd/%s/          # Application entrypoints
│   └── main.go
├── internal/        # Private application code
├── pkg/            # Public library code
├── api/            # API definitions (OpenAPI/Swagger specs, protocol definition files)
├── configs/        # Configuration files
├── scripts/        # Scripts for various build, install, analysis, etc operations
├── test/           # Additional external test apps and test data
├── docs/           # Design and user documents
├── examples/       # Examples for your applications and/or public libraries
├── go.mod
├── go.sum
├── Makefile
└── README.md
`+"```"+`

## License

This project is licensed under the MIT License - see the LICENSE file for details.
`, data.ProjectName, data.Description, data.GoVersion, data.ProjectName, data.ProjectName, data.ProjectName, data.ProjectName, data.ProjectName)
}

func GetGitignoreTemplate() string {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `+"`go test -c`"+`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Application specific
/bin/
/dist/
/tmp/
*.log
.env
.env.local
`
}

func GetMakefileTemplate(data TemplateData) string {
	return fmt.Sprintf(`# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=%s
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/%s

.PHONY: all build clean test coverage deps run help

all: test build

## build: Build the application
build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH)

## clean: Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

## test: Run tests
test:
	$(GOTEST) -v ./...

## coverage: Run tests with coverage
coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

## deps: Download and install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

## run: Run the application
run:
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH) && $(BINARY_PATH)

## dev: Run the application in development mode
dev:
	$(GOCMD) run $(MAIN_PATH)

## lint: Run linter (requires golangci-lint)
lint:
	golangci-lint run

## fmt: Format code
fmt:
	$(GOCMD) fmt ./...

## vet: Run go vet
vet:
	$(GOCMD) vet ./...

## help: Show this help message
help:
	@echo "Available targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
`, data.ProjectName, data.ProjectName)
}

// GetGoModTemplate generates go.mod content
// This is a placeholder that will be replaced by calling the main package's GenerateGoMod function
func GetGoModTemplate(data TemplateData) string {
	// Basic go.mod template - the actual implementation will call GenerateGoMod from main package
	return fmt.Sprintf(`module %s

go %s

require (
	// Dependencies will be added based on framework selection
)
`, data.ModuleName, data.GoVersion)
}