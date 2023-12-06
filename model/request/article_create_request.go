package request

type ArticleCreateRequest struct {
	Title    string `validate:"required min=1, max=255" json:"Title"`
	Body     string `validate:"required min=1, max=65,535" json:"Body"`
	Category string `validate:"required min=1, max=16" json:"Category"`
}
