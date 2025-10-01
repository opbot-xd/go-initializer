package templates

import "fmt"

// GetSimpleProjectTemplates returns templates for simple project structure
func GetSimpleProjectTemplates(data TemplateData) []Template {
	templates := GetCommonTemplates(data)
	
	templates = append(templates, []Template{
		{
			Path:    fmt.Sprintf("cmd/%s/main.go", data.ProjectName),
			Content: GetSimpleMainTemplate(data),
		},
		{
			Path:    "internal/app/app.go",
			Content: GetSimpleAppTemplate(data),
		},
		{
			Path:    fmt.Sprintf("pkg/%s/service.go", data.ProjectName),
			Content: GetSimpleServiceTemplate(data),
		},
	}...)
	
	return templates
}

func GetSimpleMainTemplate(data TemplateData) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"log"

	"%s/internal/app"
)

func main() {
	fmt.Printf("Starting %%s...\n", "%s")
	
	application := app.New()
	if err := application.Run(); err != nil {
		log.Fatalf("Application failed: %%v", err)
	}
}
`, data.ModuleName, data.ProjectName)
}

func GetSimpleAppTemplate(data TemplateData) string {
	return fmt.Sprintf(`package app

import (
	"fmt"

	"%s/pkg/%s"
)

type App struct {
	service *%s.Service
}

func New() *App {
	return &App{
		service: %s.New(),
	}
}

func (a *App) Run() error {
	fmt.Println("Application is running...")
	
	result := a.service.Process("Hello, World!")
	fmt.Printf("Service result: %%s\n", result)
	
	return nil
}
`, data.ModuleName, data.ProjectName, data.ProjectName, data.ProjectName)
}

func GetSimpleServiceTemplate(data TemplateData) string {
	return fmt.Sprintf(`package %s

import "fmt"

type Service struct {
	// Add your service dependencies here
}

func New() *Service {
	return &Service{}
}

func (s *Service) Process(input string) string {
	return fmt.Sprintf("Processed: %%s", input)
}

// Add more service methods here
`, data.ProjectName)
}