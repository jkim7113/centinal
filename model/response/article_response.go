package response

type ArticleResponse struct {
	UUID      string `json:"UUID"`
	Title     string `json:"Title"`
	Body      string `json:"Body"`
	Date      string `json:"Date"`
	ERT       int    `json:"ERT"`
	Category  string `json:"Category"`
	Thumbnail string `json:"Thumbnail"`
}
