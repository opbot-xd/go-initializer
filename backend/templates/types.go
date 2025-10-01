package templates

// Template represents a file template with its path and content
type Template struct {
	Path    string
	Content string
}

// TemplateData holds the data for template rendering
type TemplateData struct {
	ProjectName string
	ModuleName  string
	Description string
	Framework   string
	GoVersion   string
	ProjectType string
}

// GetTemplatesForProjectType returns templates based on project type and framework
func GetTemplatesForProjectType(projectType, framework string, data TemplateData) []Template {
	switch projectType {
	case "microservice":
		return GetMicroserviceTemplates(framework, data)
	case "cli-app":
		return GetCLIAppTemplates(framework, data)
	case "api-server":
		return GetAPIServerTemplates(framework, data)
	case "simple-project":
		return GetSimpleProjectTemplates(data)
	default:
		return GetSimpleProjectTemplates(data)
	}
}