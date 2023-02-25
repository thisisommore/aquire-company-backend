package getcompany

type GetCompaniesQuery struct {
	SortBy      string `form:"sortBy"`
	StartOffSet *int   `form:"startOffSet" binding:"required"`
}
