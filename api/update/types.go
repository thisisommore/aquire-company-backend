package update

type UpdateRequest struct {
	ProfileUrl        string `json:"profileUrl"`
	Name              string `json:"name"`
	OpenToAquire      *bool  `json:"openToAquire"`
	Price             int    `json:"price"`
	Product           string `json:"product"`
	CurrentYearProfit int    `json:"currentYearProfit"`
	Domain            string `json:"domain"`
}
