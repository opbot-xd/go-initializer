package main

import (
	"testing"
)

func TestGetTemplatesForProjectType(t *testing.T) {
	testCases := []struct {
		name        string
		projectType string
		framework   string
		data        TemplateData
		expectFiles int
	}{
		{
			name:        "Microservice with Gin",
			projectType: "microservice",
			framework:   "gin",
			data: TemplateData{
				ProjectName: "test-service",
				ModuleName:  "github.com/user/test-service",
				Description: "A test microservice",
				Framework:   "gin",
				GoVersion:   "1.22.0",
				ProjectType: "microservice",
			},
			expectFiles: 11, // Common files + microservice specific files
		},
		{
			name:        "CLI App with Cobra",
			projectType: "cli-app",
			framework:   "cobra",
			data: TemplateData{
				ProjectName: "test-cli",
				ModuleName:  "github.com/user/test-cli",
				Description: "A test CLI application",
				Framework:   "cobra",
				GoVersion:   "1.22.0",
				ProjectType: "cli-app",
			},
			expectFiles: 7, // Common files + CLI specific files
		},
		{
			name:        "Simple Project",
			projectType: "simple-project",
			framework:   "golly",
			data: TemplateData{
				ProjectName: "test-simple",
				ModuleName:  "github.com/user/test-simple",
				Description: "A simple test project",
				Framework:   "golly",
				GoVersion:   "1.22.0",
				ProjectType: "simple-project",
			},
			expectFiles: 7, // Common files + simple project files
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			templates := GetTemplatesForProjectType(tc.projectType, tc.framework, tc.data)
			
			if len(templates) != tc.expectFiles {
				t.Errorf("Expected %d templates, got %d", tc.expectFiles, len(templates))
			}
			
			// Check that all templates have content
			for _, template := range templates {
				if template.Path == "" {
					t.Errorf("Template has empty path")
				}
				if template.Content == "" {
					t.Errorf("Template %s has empty content", template.Path)
				}
			}
		})
	}
}

func TestTemplateGeneration(t *testing.T) {
	request := CreateProjectRequest{
		ProjectType: "microservice",
		GoVersion:   "1.22.0",
		Framework:   "gin",
		ModuleName:  "github.com/user/test-service",
		Name:        "test-service",
		Description: "A test microservice for template validation",
	}

	buf, err := GenerateProjectFromTemplate(request)
	if err != nil {
		t.Fatalf("Failed to generate project: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated zip file is empty")
	}
}