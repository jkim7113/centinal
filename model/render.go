package model

import "github.com/jkim7113/centinal/model/response"

type DataToRender struct {
	Data     []response.ArticleResponse
	Path string
}