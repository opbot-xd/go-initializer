package templates

// GetAPIServerTemplates returns templates for API server projects
func GetAPIServerTemplates(framework string, data TemplateData) []Template {
	// API Server is similar to microservice but with more focus on REST API
	templates := GetMicroserviceTemplates(framework, data)
	
	// Add additional API-specific templates
	templates = append(templates, []Template{
		{
			Path:    "internal/middleware/cors.go",
			Content: GetCORSMiddlewareTemplate(framework),
		},
		{
			Path:    "internal/middleware/auth.go",
			Content: GetAuthMiddlewareTemplate(framework),
		},
	}...)
	
	return templates
}

func GetCORSMiddlewareTemplate(framework string) string {
	switch framework {
	case "gin":
		return `package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}
`
	case "echo":
		return `package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	})
}
`
	default:
		return `package middleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
`
	}
}

func GetAuthMiddlewareTemplate(framework string) string {
	switch framework {
	case "gin":
		return `package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Basic token validation (implement your own logic)
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !isValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	})
}

func isValidToken(token string) bool {
	// Implement your token validation logic here
	// This is a placeholder implementation
	return token != ""
}
`
	case "echo":
		return `package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header required"})
			}

			// Basic token validation (implement your own logic)
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authorization header format"})
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if !isValidToken(token) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			return next(c)
		}
	}
}

func isValidToken(token string) bool {
	// Implement your token validation logic here
	// This is a placeholder implementation
	return token != ""
}
`
	default:
		return `package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Authorization header required"})
			return
		}

		// Basic token validation (implement your own logic)
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid authorization header format"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !isValidToken(token) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	// Implement your token validation logic here
	// This is a placeholder implementation
	return token != ""
}
`
	}
}