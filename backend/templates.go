package main

import (
	"./templates"
)

// Re-export types and functions from templates package
type Template = templates.Template
type TemplateData = templates.TemplateData

// GetTemplatesForProjectType returns templates based on project type and framework
func GetTemplatesForProjectType(projectType, framework string, data TemplateData) []Template {
	// Get templates from the templates package
	templateList := templates.GetTemplatesForProjectType(projectType, framework, data)
	
	// Replace the go.mod template with the proper one that uses our existing GenerateGoMod function
	for i, template := range templateList {
		if template.Path == "go.mod" {
			templateList[i].Content = GenerateGoMod(data.ModuleName, data.GoVersion, data.Framework, data.ProjectType)
		}
	}
	
	return templateList
}