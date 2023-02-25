// Package api provide support to create /api group
package api

import (
	"template-app/api/auth"
	"template-app/api/publish"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /api group and v1 routes to given gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth.ApplyRoutes(api)
		publish.ApplyRoutes(api)
	}
}
