package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TemplatePreviewRequest represents a request to preview templates
type TemplatePreviewRequest struct {
	ProjectType string `json:"projectType"`
	Framework   string `json:"framework"`
	ProjectName string `json:"projectName"`
	ModuleName  string `json:"moduleName"`
	Description string `json:"description"`
	GoVersion   string `json:"goVersion"`
}

// TemplatePreviewResponse represents the response with template previews
type TemplatePreviewResponse struct {
	Templates []Template `json:"templates"`
	Count     int        `json:"count"`
}

// PreviewTemplatesHandler handles template preview requests
func PreviewTemplatesHandler(ctx *gin.Context) {
	var request TemplatePreviewRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("[ERROR] Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("[INFO] Received template preview request: %+v", request)

	// Create template data
	templateData := TemplateData{
		ProjectName: request.ProjectName,
		ModuleName:  request.ModuleName,
		Description: request.Description,
		Framework:   request.Framework,
		GoVersion:   request.GoVersion,
		ProjectType: request.ProjectType,
	}

	// Generate templates
	templates := GetTemplatesForProjectType(request.ProjectType, request.Framework, templateData)

	response := TemplatePreviewResponse{
		Templates: templates,
		Count:     len(templates),
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAvailableTemplatesHandler returns available template types and frameworks
func GetAvailableTemplatesHandler(ctx *gin.Context) {
	response := map[string]interface{}{
		"projectTypes": SupportedProjectTypesLabelsMap,
		"frameworks":   SupportedFrameworksMap,
		"goVersions":   SupportedGoVersionsMap,
	}

	ctx.JSON(http.StatusOK, response)
}

// TemplateStatsHandler returns statistics about templates
func TemplateStatsHandler(ctx *gin.Context) {
	projectType := ctx.Query("projectType")
	framework := ctx.Query("framework")

	if projectType == "" {
		projectType = "microservice"
	}
	if framework == "" {
		framework = "gin"
	}

	// Create template data for stats
	templateData := TemplateData{
		ProjectName: "project",
		ModuleName:  "github.com/user/project",
		Description: "Project for statistics",
		Framework:   framework,
		GoVersion:   "1.22.0",
		ProjectType: projectType,
	}

	templates := GetTemplatesForProjectType(projectType, framework, templateData)

	// Calculate statistics
	stats := map[string]interface{}{
		"totalFiles":   len(templates),
		"goFiles":      countFilesByExtension(templates, ".go"),
		"configFiles":  countConfigFiles(templates),
		"totalLines":   countTotalLines(templates),
		"projectType":  projectType,
		"framework":    framework,
	}

	ctx.JSON(http.StatusOK, stats)
}

func countFilesByExtension(templates []Template, extension string) int {
	count := 0
	for _, template := range templates {
		if len(template.Path) >= len(extension) && 
		   template.Path[len(template.Path)-len(extension):] == extension {
			count++
		}
	}
	return count
}

func countConfigFiles(templates []Template) int {
	count := 0
	for _, template := range templates {
		if template.Path == "go.mod" || 
		   template.Path == "Makefile" || 
		   template.Path == ".gitignore" ||
		   len(template.Path) >= 5 && template.Path[len(template.Path)-5:] == ".yaml" ||
		   len(template.Path) >= 4 && template.Path[len(template.Path)-4:] == ".yml" {
			count++
		}
	}
	return count
}

func countTotalLines(templates []Template) int {
	total := 0
	for _, template := range templates {
		lines := 1
		for _, char := range template.Content {
			if char == '\n' {
				lines++
			}
		}
		total += lines
	}
	return total
}