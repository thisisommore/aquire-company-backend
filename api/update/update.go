package update

import (
	"net/http"
	pasetomiddleware "template-app/api/middleware/auth/paseto"
	"template-app/models/company"
	"template-app/models/company/companymethods"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/update")
	{
		g.PATCH("", update)
	}
}

func update(c *gin.Context) {
	var body UpdateRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}
	emailId := c.GetString(pasetomiddleware.EmailIdInContext)
	if err = companymethods.Update(emailId, company.Company{
		Name:                 body.Name,
		OpenToAquire:         body.OpenToAquire,
		Price:                body.Price,
		Product:              body.Product,
		CurrentYearProfit:    body.CurrentYearProfit,
		Domain:               body.Domain,
		ProfileUrl:           body.ProfileUrl,
		Ceo:                  body.Ceo,
		Cto:                  body.Cto,
		Owner:                body.Owner,
		Address:              body.Address,
		Description:          body.Description,
		NetWorth:             body.NetWorth,
		LinkedIn:             body.LinkedIn,
		Website:              body.Website,
		Youtube:              body.Youtube,
		BuisinessLicenseLink: body.BuisinessLicenseLink,
		LiveUsersLink:        body.LiveUsersLink,
	}); err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}
	httpo.NewSuccessResponse(http.StatusOK, "data updated").
		SendD(c)
}
