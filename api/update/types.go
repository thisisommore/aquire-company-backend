package update

type UpdateRequest struct {
	Name         string `json:"name"`
	OpenToAquire *bool  `json:"openToAquire"`
	Price        int    `json:"price"`
}
