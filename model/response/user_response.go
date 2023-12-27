package response

type UserResponse struct {
	UUID     string `json:"UUID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Bio      string `json:"Bio"`
	Date     string `json:"Date"`
	PFP      string `json:"PFP"`
	Role     string `json:"Role"`
}
