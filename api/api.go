// Package api provide support to create /api group
package api

import (
	"template-app/api/auth"
	"template-app/api/getcompany"
	"template-app/api/getmycompany"
	pasetomiddleware "template-app/api/middleware/auth/paseto"
	"template-app/api/update"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /api group and v1 routes to given gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth.ApplyRoutes(api)
		getcompany.ApplyRoutes(api)
		api.Use(pasetomiddleware.PASETO)
		update.ApplyRoutes(api)
		getmycompany.ApplyRoutes(api)
	}
}
