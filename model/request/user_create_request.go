package request

type UserCreateRequest struct {
	Username string `validate:"required min=3, max=24" json:"Username"`
	Email    string `validate:"required min=1, max=65,535" json:"Email"`
	Pw       string `validate:"required min=8, max=255" json:"Pw"`
	Bio      string `json:"Bio"`
	PFP      string `json:"PFP"`
}
