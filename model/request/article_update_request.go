package request

type ArticleUpdateRequest struct {
	Id       int    `json:"Id"`
	Title    string `validate:"required min=1, max=255" json:"Title"`
	Body     string `validate:"required min=1, max=65,535" json:"Body"`
	ERT      int    `json:"ERT"`
	Category string `validate:"required min=1, max=16" json:"Category"`
}
