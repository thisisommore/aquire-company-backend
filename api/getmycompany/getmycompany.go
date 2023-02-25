package getmycompany

import (
	"net/http"
	pasetomiddleware "template-app/api/middleware/auth/paseto"
	"template-app/models/company/companymethods"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/get-my-company")
	{
		g.GET("", getMyCompany)
	}
}

func getMyCompany(c *gin.Context) {

	emailId := c.GetString(pasetomiddleware.EmailIdInContext)
	companyDetails, err := companymethods.Get(emailId)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}
	httpo.NewSuccessResponseP(http.StatusOK, "details fetched", companyDetails).
		SendD(c)
}
