package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
)

func GenerateMicroservice(request CreateProjectRequest) (*bytes.Buffer, error) {
	return GenerateProjectFromTemplate(request)
}

func GenerateSimpleProject(request CreateProjectRequest) (*bytes.Buffer, error) {
	return GenerateProjectFromTemplate(request)
}

// GenerateProjectFromTemplate generates a project using the template system
func GenerateProjectFromTemplate(request CreateProjectRequest) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	folderName := request.Name
	if folderName == "" {
		folderName = "project"
	}

	// Prepare template data
	templateData := TemplateData{
		ProjectName: folderName,
		ModuleName:  request.ModuleName,
		Description: request.Description,
		Framework:   strings.ToLower(request.Framework),
		GoVersion:   request.GoVersion,
		ProjectType: request.ProjectType,
	}

	// Get templates for the project type
	templates := GetTemplatesForProjectType(request.ProjectType, strings.ToLower(request.Framework), templateData)

	// Generate files from templates
	for _, template := range templates {
		filePath := fmt.Sprintf("%s/%s", folderName, template.Path)
		
		file, err := zipWriter.Create(filePath)
		if err != nil {
			log.Printf("[ERROR] Failed to create file %s in zip: %v", filePath, err)
			return nil, fmt.Errorf("failed to create file %s in zip", template.Path)
		}
		
		_, err = file.Write([]byte(template.Content))
		if err != nil {
			log.Printf("[ERROR] Failed to write file %s: %v", filePath, err)
			return nil, fmt.Errorf("failed to write file %s", template.Path)
		}
		
		log.Printf("[INFO] Generated file: %s", filePath)
	}

	err := zipWriter.Close()
	if err != nil {
		log.Printf("[ERROR] Failed to close zip writer: %v", err)
		return nil, errors.New("failed to finalize zip file")
	}

	log.Printf("[INFO] Successfully generated %s project with %d files", request.ProjectType, len(templates))
	return buf, nil
}
