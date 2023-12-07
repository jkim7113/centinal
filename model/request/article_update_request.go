package request

type ArticleUpdateRequest struct {
	UUID      string `json:"UUID"`
	Title     string `validate:"required min=1, max=255" json:"Title"`
	Body      string `validate:"required min=1, max=65,535" json:"Body"`
	ERT       int    `json:"ERT"`
	Category  string `validate:"required min=1, max=16" json:"Category"`
	Thumbnail string `json:"Thumbnail"`
}
