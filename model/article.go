package model

import "github.com/jkim7113/centinal/model/response"

type Article struct {
	UUID      string
	Title     string
	Body      string
	Author    response.UserResponse
	Date      string
	ERT       int
	Category  string
	Thumbnail string
}
