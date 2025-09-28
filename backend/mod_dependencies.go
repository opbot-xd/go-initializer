package main

import (
	"fmt"
)

func GenerateGoMod(moduleName, goVersion, framework, projectType string) string {
	var requiredImports []string

	// Validate project type
	if !SupportedProjectTypesMap[projectType] {
		return fmt.Sprintf("// Unsupported project type: %s\n", projectType)
	}

	// Validate framework for the project type
	frameworks, ok := SupportedFrameworksMap[projectType]
	if !ok || !frameworks[framework] {
		return fmt.Sprintf("// Unsupported framework '%s' for project type '%s'\n", framework, projectType)
	}

	// Get dependencies for the framework
	if deps, ok := DependencyMap[framework]; ok {
		requiredImports = deps
	} else {
		requiredImports = []string{}
	}

	goMod := fmt.Sprintf("module %s\n\ngo %s\n", moduleName, goVersion)
	if len(requiredImports) > 0 {
		goMod += "\nrequire (\n"
		for _, imp := range requiredImports {
			goMod += fmt.Sprintf("\t%s latest\n", imp)
		}
		goMod += ")\n"
	}
	return goMod
}
