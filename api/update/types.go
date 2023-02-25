package update

type UpdateRequest struct {
	ProfileUrl           string `json:"profileUrl"`
	Name                 string `json:"name"`
	OpenToAquire         *bool  `json:"openToAquire"`
	Price                int    `json:"price"`
	Product              string `json:"product"`
	CurrentYearProfit    int    `json:"currentYearProfit"`
	Domain               string `json:"domain"`
	Ceo                  string `json:"ceo"`
	Owner                string `json:"owner"`
	Cto                  string `json:"cto"`
	Address              string `json:"address"`
	Description          string `json:"Description"`
	NetWorth             int    `json:"networth"`
	LinkedIn             string `json:"linkedIn"`
	LiveUsersLink        string `json:"liveUserLink"`
	Website              string `json:"website"`
	Youtube              string `json:"youtube"`
	BuisinessLicenseLink string `json:"businessLicenseLink"`
	ActiveUsers          int    `json:"activeUsers"`
}
