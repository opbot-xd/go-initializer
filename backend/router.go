package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewRouter(v *validator.Validate) *gin.Engine {
	service := gin.Default()

	// Configure CORS with more permissive settings for development
	service.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Common React dev server ports
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	service.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "server is up and running")
	})

	// API routes
	api := service.Group("/api")
	{
		api.GET("/meta", MetaHandler)
		api.POST("/generate", GenerateHandler)
		
		// Template management routes
		api.POST("/preview", PreviewTemplatesHandler)
		api.GET("/templates", GetAvailableTemplatesHandler)
		api.GET("/templates/stats", TemplateStatsHandler)
	}

	// Legacy routes for backward compatibility
	service.GET("/meta", MetaHandler)
	service.POST("/generate", GenerateHandler)

	// Serve static files from frontend build
	service.Static("/static", "./frontend/build/static")
	service.StaticFile("/", "./frontend/build/index.html")
	service.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
	
	// Catch-all route for React Router (SPA)
	service.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	return service
}
