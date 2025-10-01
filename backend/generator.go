package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"log"
)

func GenerateMicroservice(request CreateProjectRequest) (*bytes.Buffer, error) {
	// Implementation for generating a microservice project based on input
	// generate a zip file in memory

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Create a folder with the project name and write files inside it
	folderName := request.Name
	if folderName == "" {
		folderName = "project"
	}

	// Example: add a README.md file inside the folder
	readmeContent := fmt.Sprintf("# %s\n\n%s", request.Name, request.Description)
	readmeFile, err := zipWriter.Create(fmt.Sprintf("%s/README.md", folderName))
	if err != nil {
		log.Printf("[ERROR] Failed to create file in zip: %v", err)
		err = errors.New("failed to create file in zip")
		return nil, err
	}
	_, err = readmeFile.Write([]byte(readmeContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write to zip: %v", err)
		err = errors.New("failed to write to zip file")
		return nil, err
	}

	// Add more files as needed based on request
	err = zipWriter.Close()
	if err != nil {
		log.Printf("[ERROR] Failed to close zip writer: %v", err)
		err = errors.New("failed to finalize zip file")
		return nil, err
	}

	return buf, nil
}

// myproject/
// ├── cmd/
// │   └── myproject/
// │       └── main.go
// ├── internal/           # Business logic
// │   └── service.go
// ├── go.mod
// └── README.md
func GenerateSimpleProjecet(request CreateProjectRequest) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	folderName := request.Name
	if folderName == "" {
		folderName = "myproject"
	}

	// 1. README.md
	readmeContent := fmt.Sprintf("# %s\n\n%s", folderName, request.Description)
	readmeFile, err := zipWriter.Create(fmt.Sprintf("%s/README.md", folderName))
	if err != nil {
		log.Printf("[ERROR] Failed to create README.md in zip: %v", err)
		return nil, errors.New("failed to create README.md in zip")
	}
	_, err = readmeFile.Write([]byte(readmeContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write README.md: %v", err)
		return nil, errors.New("failed to write README.md")
	}

	// 2. go.mod (minimal content)
	gomodContent := GenerateGoMod(folderName, request.GoVersion, request.Framework, request.ProjectType)
	fmt.Println("Generated go.mod content:\n", gomodContent) // Debug log
	gomodFile, err := zipWriter.Create(fmt.Sprintf("%s/go.mod", folderName))
	if err != nil {
		log.Printf("[ERROR] Failed to create go.mod in zip: %v", err)
		return nil, errors.New("failed to create go.mod in zip")
	}
	_, err = gomodFile.Write([]byte(gomodContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write go.mod: %v", err)
		return nil, errors.New("failed to write go.mod")
	}

	// 3. cmd/myproject/main.go
	mainGoContent := "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello from main!\")\n}\n"
	mainGoPath := fmt.Sprintf("%s/cmd/%s/main.go", folderName, folderName)
	mainGoFile, err := zipWriter.Create(mainGoPath)
	if err != nil {
		log.Printf("[ERROR] Failed to create main.go in zip: %v", err)
		return nil, errors.New("failed to create main.go in zip")
	}
	_, err = mainGoFile.Write([]byte(mainGoContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write main.go: %v", err)
		return nil, errors.New("failed to write main.go")
	}

	// 4. internal/service.go
	serviceGoContent := "package internal\n\n// Business logic goes here\nfunc Service() string {\n    return \"Service logic\"\n}\n"
	serviceGoPath := fmt.Sprintf("%s/internal/service.go", folderName)
	serviceGoFile, err := zipWriter.Create(serviceGoPath)
	if err != nil {
		log.Printf("[ERROR] Failed to create service.go in zip: %v", err)
		return nil, errors.New("failed to create service.go in zip")
	}
	_, err = serviceGoFile.Write([]byte(serviceGoContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write service.go: %v", err)
		return nil, errors.New("failed to write service.go")
	}

	err = zipWriter.Close()
	if err != nil {
		log.Printf("[ERROR] Failed to close zip writer: %v", err)
		return nil, errors.New("failed to finalize zip file")
	}

	return buf, nil
}
