package request

type UserUpdateRequest struct {
	UUID     string `json:"UUID"`
	Username string `validate:"required min=3, max=24" json:"Username"`
	Email    string `validate:"required min=1, max=65,535" json:"Email"`
	Bio      string `json:"Bio"`
	PFP      string `json:"PFP"`
}
