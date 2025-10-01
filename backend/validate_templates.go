package main

import (
	"fmt"
	"log"
)

// Simple validation function to check template generation
func validateTemplates() {
	testData := TemplateData{
		ProjectName: "test-project",
		ModuleName:  "github.com/user/test-project",
		Description: "A test project for validation",
		Framework:   "gin",
		GoVersion:   "1.22.0",
		ProjectType: "microservice",
	}

	// Test each project type
	projectTypes := []string{"microservice", "cli-app", "api-server", "simple-project"}
	frameworks := map[string][]string{
		"microservice": {"gin", "echo", "fiber"},
		"cli-app":      {"cobra", "urfave"},
		"api-server":   {"gin", "echo", "chi"},
		"simple-project": {"golly"},
	}

	for _, projectType := range projectTypes {
		testData.ProjectType = projectType
		for _, framework := range frameworks[projectType] {
			testData.Framework = framework
			
			templates := GetTemplatesForProjectType(projectType, framework, testData)
			
			fmt.Printf("Project Type: %s, Framework: %s, Templates: %d\n", 
				projectType, framework, len(templates))
			
			// Check that we have essential files
			hasReadme := false
			hasGoMod := false
			hasMain := false
			
			for _, template := range templates {
				switch template.Path {
				case "README.md":
					hasReadme = true
				case "go.mod":
					hasGoMod = true
				}
				if template.Path == fmt.Sprintf("cmd/%s/main.go", testData.ProjectName) {
					hasMain = true
				}
			}
			
			if !hasReadme || !hasGoMod || !hasMain {
				log.Printf("Missing essential files for %s/%s: README=%v, go.mod=%v, main.go=%v",
					projectType, framework, hasReadme, hasGoMod, hasMain)
			}
		}
	}
	
	fmt.Println("Template validation completed!")
}

func main() {
	validateTemplates()
}