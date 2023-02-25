package getcompany

import (
	"fmt"
	"net/http"
	"template-app/models/company/companymethods"
	"template-app/pkg/httphelper"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/get-companies")
	{
		g.GET("", GetCompany)
	}
}

func GetCompany(c *gin.Context) {
	var reqQuery GetCompaniesQuery
	err := c.ShouldBindQuery(&reqQuery)
	if err != nil {
		err := fmt.Errorf("body is invalid: %w", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}
	sortBy := companymethods.SORT_BY_NONE
	switch reqQuery.SortBy {
	case "price":
		sortBy = companymethods.SORT_BY_PRICE
	case "users":
		sortBy = companymethods.SORT_BY_USERS
	case "current_year_profite":
		sortBy = companymethods.SORT_BY_CURRENT_YEAR_PROFIT
	}

	companies, err := companymethods.GetCompaniesOpenToAquire(*reqQuery.StartOffSet, 10, sortBy)
	if err != nil {
		httphelper.HandleDBError(err, "failed to fetch companies", c)
		return
	}
	httpo.NewSuccessResponseP(200, "companies fetched", companies).SendD(c)
}
