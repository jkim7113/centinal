package controller

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/model/request"
	"github.com/jkim7113/centinal/model/response"
	"github.com/jkim7113/centinal/service"
	"github.com/jkim7113/centinal/util"
)

type ArticleController struct {
	ArticleService service.ArticleService
}

type dataToRender struct {
	Data []response.ArticleResponse
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
	UUID := chi.URLParam(r, "UUID")

	result := controller.ArticleService.FindById(r.Context(), UUID)
	HTTPResponse := response.HTTPResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	result := controller.ArticleService.FindAll(r.Context())
	tmpl := template.Must(template.ParseFiles("./views/index.html", "./views/config.tmpl", "./views/header.tmpl"))
	tmpl.Execute(w, dataToRender{Data: result})
}

func (controller *ArticleController) FindByCategory(w http.ResponseWriter, r *http.Request) {
	Category := chi.URLParam(r, "Category")

	result := controller.ArticleService.FindByCategory(r.Context(), Category)
	tmpl := template.Must(template.ParseFiles("./views/category.html", "./views/config.tmpl", "./views/header.tmpl"))
	tmpl.Execute(w, dataToRender{Data: result})
}
