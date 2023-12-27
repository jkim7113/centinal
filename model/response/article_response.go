package response

type ArticleResponse struct {
	UUID       string `json:"UUID"`
	Title      string `json:"Title"`
	Body       string `json:"Body"`
	Author     string `json:"Author"`
	AuthorPFP  string `json:"AuthorPFP"`
	AuthorUUID string `json:"AuthorUUID"`
	Date       string `json:"Date"`
	ERT        int    `json:"ERT"`
	Category   string `json:"Category"`
	Thumbnail  string `json:"Thumbnail"`
}
