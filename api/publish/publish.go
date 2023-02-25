package publish

import (
	"net/http"
	"template-app/models/sitemodel"
	"template-app/models/sitemodel/site_model_methods"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/publish")
	{
		g.POST("", publish)
	}
}

func publish(c *gin.Context) {
	var body PublishRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}

	if err = site_model_methods.Add(&sitemodel.Site{Name: body.Name}); err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}
	httpo.NewSuccessResponse(http.StatusOK, "site deployed successfully").
		SendD(c)

}
