package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateHandler(ctx *gin.Context) {
	var request CreateProjectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("[ERROR] Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Printf("[INFO] Received request: %+v", request)

	// generate a zip file in memory

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Create a folder with the project name and write files inside it
	folderName := request.Name
	if folderName == "" {
		folderName = "project"
	}

	// Example: add a README.md file inside the folder
	readmeContent := fmt.Sprintf("# %s\n\n%s", request.Name, request.Description)
	readmeFile, err := zipWriter.Create(fmt.Sprintf("%s/README.md", folderName))
	if err != nil {
		log.Printf("[ERROR] Failed to create file in zip: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create zip file"})
		return
	}
	_, err = readmeFile.Write([]byte(readmeContent))
	if err != nil {
		log.Printf("[ERROR] Failed to write to zip: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to zip file"})
		return
	}

	// Add more files as needed based on request
	err = zipWriter.Close()
	if err != nil {
		log.Printf("[ERROR] Failed to close zip writer: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize zip file"})
		return
	}

	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Disposition", "attachment; filename=project.zip")
	ctx.Data(http.StatusOK, "application/zip", buf.Bytes())
}
