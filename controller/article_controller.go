package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/model/request"
	"github.com/jkim7113/centinal/model/response"
	"github.com/jkim7113/centinal/service"
	"github.com/jkim7113/centinal/util"
)

type ArticleController struct {
	ArticleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) *ArticleController {
	return &ArticleController{ArticleService: articleService}
}

func (controller *ArticleController) Create(w http.ResponseWriter, r *http.Request) {
	articleCreateRequest := request.ArticleCreateRequest{}
	util.DecodeRequestBody(r, &articleCreateRequest)

	controller.ArticleService.Create(r.Context(), articleCreateRequest)
	HTTPResponse := response.HTTPResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *ArticleController) FindById(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "Id")
	Id, err := strconv.Atoi(param)
	util.PanicIfError(err)

	result := controller.ArticleService.FindById(r.Context(), Id)
	HTTPResponse := response.HTTPResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}
