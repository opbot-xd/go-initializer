package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetaHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"supportedProjectTypes": SupportedProjectTypesLabelsMap,
		"supportedGoVersions":   SupportedGoVersionsMap,
		"supportedFrameworks":   SupportedFrameworksMap,
	})
}

func GenerateHandler(ctx *gin.Context) {
	var request CreateProjectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("[ERROR] Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	// Validate request
	if !SupportedProjectTypesMap[request.ProjectType] {
		log.Printf("[ERROR] Unsupported project type: %s", request.ProjectType)
		resp := ErrorResponseBody{
			StatusCode: http.StatusBadRequest,
			Message:    "Unsupported project type",
		}
		resp.GenerateResponse(ctx)
		return
	}
	
	log.Printf("[INFO] Received request: %+v", request)
	
	// Generate project using template system
	buf, err := GenerateProjectFromTemplate(request)
	if err != nil {
		log.Printf("[ERROR] Failed to generate project: %v", err)
		resp := ErrorResponseBody{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to generate project",
		}
		resp.GenerateResponse(ctx)
		return
	}
	
	resp := SuccessResponseBody{Data: buf.Bytes()}
	resp.GenerateResponse(ctx)
}
