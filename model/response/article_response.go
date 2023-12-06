package response

type ArticleResponse struct {
	Id       int    `json:"Id"`
	Title    string `json:"Title"`
	Body     string `json:"Body"`
	Date     string `json:"Date"`
	ERT      int    `json:"ERT"`
	Category string
}
