package publish

type PublishRequest struct {
	Name       string `json:"name" binding:"required"`
	HtmlString string `json:"htmlString"`
	Components []string `json:"components"`
}
