package main

type CreateProjectRequest struct {
	ProjectType string `json:"projectType"`
	GoVersion   string `json:"goVersion"`
	Framework   string `json:"framework"`
	ModuleName  string `json:"moduleName"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
