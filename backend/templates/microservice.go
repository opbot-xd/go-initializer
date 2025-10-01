package templates

import "fmt"

// GetMicroserviceTemplates returns templates for microservice projects
func GetMicroserviceTemplates(framework string, data TemplateData) []Template {
	templates := GetCommonTemplates(data)
	
	// Add microservice-specific templates
	templates = append(templates, []Template{
		{
			Path:    fmt.Sprintf("cmd/%s/main.go", data.ProjectName),
			Content: GetMicroserviceMainTemplate(framework, data),
		},
		{
			Path:    "internal/server/server.go",
			Content: GetMicroserviceServerTemplate(framework, data),
		},
		{
			Path:    "internal/handler/health.go",
			Content: GetHealthHandlerTemplate(framework, data),
		},
		{
			Path:    "internal/handler/handler.go",
			Content: GetHandlerTemplate(framework, data),
		},
		{
			Path:    "internal/config/config.go",
			Content: GetConfigTemplate(data),
		},
		{
			Path:    "configs/config.yaml",
			Content: GetConfigYamlTemplate(data),
		},
		{
			Path:    "api/openapi.yaml",
			Content: GetOpenAPITemplate(data),
		},
	}...)
	
	return templates
}

func GetMicroserviceMainTemplate(framework string, data TemplateData) string {
	switch framework {
	case "gin":
		return fmt.Sprintf(`package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"%s/internal/config"
	"%s/internal/handler"
	"%s/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %%v", err)
	}

	// Set gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize handlers
	h := handler.New()

	// Initialize and start server
	srv := server.New(cfg, h)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %%v", err)
	}
}
`, data.ModuleName, data.ModuleName, data.ModuleName)
	case "echo":
		return fmt.Sprintf(`package main

import (
	"log"

	"%s/internal/config"
	"%s/internal/handler"
	"%s/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %%v", err)
	}

	// Initialize handlers
	h := handler.New()

	// Initialize and start server
	srv := server.New(cfg, h)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %%v", err)
	}
}
`, data.ModuleName, data.ModuleName, data.ModuleName)
	default:
		return fmt.Sprintf(`package main

import (
	"fmt"
	"log"
	"net/http"

	"%s/internal/config"
	"%s/internal/handler"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %%v", err)
	}

	// Initialize handlers
	h := handler.New()

	// Setup routes
	http.HandleFunc("/health", h.Health)
	http.HandleFunc("/", h.Hello)

	// Start server
	addr := fmt.Sprintf(":%%s", cfg.Port)
	log.Printf("Server starting on %%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %%v", err)
	}
}
`, data.ModuleName, data.ModuleName)
	}
}

func GetMicroserviceServerTemplate(framework string, data TemplateData) string {
	switch framework {
	case "gin":
		return fmt.Sprintf(`package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"%s/internal/config"
	"%s/internal/handler"
)

type Server struct {
	config  *config.Config
	handler *handler.Handler
	router  *gin.Engine
}

func New(cfg *config.Config, h *handler.Handler) *Server {
	return &Server{
		config:  cfg,
		handler: h,
		router:  gin.New(),
	}
}

func (s *Server) Start() error {
	// Setup middleware
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	// Setup routes
	s.setupRoutes()

	// Create HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%%s", s.config.Port),
		Handler: s.router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %%s", s.config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %%v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %%v", err)
	}

	log.Println("Server exited")
	return nil
}

func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.handler.Health)
	
	// API routes
	api := s.router.Group("/api/v1")
	{
		api.GET("/hello", s.handler.Hello)
	}
}
`, data.ModuleName, data.ModuleName)
	case "echo":
		return fmt.Sprintf(`package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"%s/internal/config"
	"%s/internal/handler"
)

type Server struct {
	config  *config.Config
	handler *handler.Handler
	echo    *echo.Echo
}

func New(cfg *config.Config, h *handler.Handler) *Server {
	return &Server{
		config:  cfg,
		handler: h,
		echo:    echo.New(),
	}
}

func (s *Server) Start() error {
	// Setup middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	// Setup routes
	s.setupRoutes()

	// Start server in a goroutine
	go func() {
		addr := fmt.Sprintf(":%%s", s.config.Port)
		log.Printf("Server starting on port %%s", s.config.Port)
		if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %%v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %%v", err)
	}

	log.Println("Server exited")
	return nil
}

func (s *Server) setupRoutes() {
	// Health check
	s.echo.GET("/health", s.handler.Health)
	
	// API routes
	api := s.echo.Group("/api/v1")
	api.GET("/hello", s.handler.Hello)
}
`, data.ModuleName, data.ModuleName)
	default:
		return `package server

// Basic HTTP server implementation
// This is a placeholder for non-framework implementations
`
	}
}

func GetHealthHandlerTemplate(framework string, data TemplateData) string {
	switch framework {
	case "gin":
		return `package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "` + data.ProjectName + `",
	})
}
`
	case "echo":
		return `package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "` + data.ProjectName + `",
	})
}
`
	default:
		return `package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "` + data.ProjectName + `",
	}
	
	json.NewEncoder(w).Encode(response)
}
`
	}
}

func GetHandlerTemplate(framework string, data TemplateData) string {
	switch framework {
	case "gin":
		return `package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// Add your dependencies here (database, services, etc.)
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from ` + data.ProjectName + `!",
	})
}
`
	case "echo":
		return `package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	// Add your dependencies here (database, services, etc.)
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello from ` + data.ProjectName + `!",
	})
}
`
	default:
		return `package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	// Add your dependencies here (database, services, etc.)
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	response := map[string]string{
		"message": "Hello from ` + data.ProjectName + `!",
	}
	
	json.NewEncoder(w).Encode(response)
}
`
	}
}

func GetConfigTemplate(data TemplateData) string {
	return `package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	Environment string
	LogLevel    string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
`
}

func GetConfigYamlTemplate(data TemplateData) string {
	return `# Configuration for ` + data.ProjectName + `
server:
  port: 8080
  environment: development

logging:
  level: info
  format: json

# Add your application-specific configuration here
`
}

func GetOpenAPITemplate(data TemplateData) string {
	return `openapi: 3.0.3
info:
  title: ` + data.ProjectName + ` API
  description: ` + data.Description + `
  version: 1.0.0
  contact:
    name: API Support
    email: support@company.com

servers:
  - url: http://localhost:8080
    description: Development server

paths:
  /health:
    get:
      summary: Health check endpoint
      description: Returns the health status of the service
      responses:
        '200':
          description: Service is healthy
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: healthy
                  timestamp:
                    type: string
                    format: date-time
                  service:
                    type: string
                    example: ` + data.ProjectName + `

  /api/v1/hello:
    get:
      summary: Hello endpoint
      description: Returns a greeting message
      responses:
        '200':
          description: Successful response
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: Hello from ` + data.ProjectName + `!

components:
  schemas:
    Error:
      type: object
      properties:
        error:
          type: string
          description: Error message
        code:
          type: integer
          description: Error code
      required:
        - error
        - code
`
}