package controller

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/model"
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
	util.DecodeFormData(r, &articleCreateRequest)

	controller.ArticleService.Create(r.Context(), articleCreateRequest)
	HTTPResponse := response.HTTPResponse{
		Code:   201,
		Status: "Created",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *ArticleController) Update(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")
	articleUpdateRequest := request.ArticleUpdateRequest{}
	articleUpdateRequest.UUID = UUID
	util.DecodeRequestBody(r, &articleUpdateRequest)

	controller.ArticleService.Update(r.Context(), articleUpdateRequest)
	HTTPResponse := response.HTTPResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *ArticleController) Delete(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")

	controller.ArticleService.Delete(r.Context(), UUID)
	HTTPResponse := response.HTTPResponse{
		Code:   204,
		Status: "No Content",
		Data:   nil,
	}
	util.EncodeResponseBody(w, HTTPResponse)
}

func (controller *ArticleController) FindById(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "UUID")

	result := controller.ArticleService.FindById(r.Context(), UUID)
	tmpl := template.Must(template.ParseFiles("./view/article.html", "./view/config.tmpl", "./view/header.tmpl", "./view/comment.tmpl", "./view/footer.tmpl"))
	tmpl.Execute(w, model.DataToRender{Data: []response.ArticleResponse{result}, Path: "/category/" + result.Category})
}

func (controller *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	result := controller.ArticleService.FindAll(r.Context())
	tmpl := template.Must(template.ParseFiles("./view/index.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
	tmpl.Execute(w, model.DataToRender{Data: result, Path: ""})
}

func (controller *ArticleController) FindByCategory(w http.ResponseWriter, r *http.Request) {
	Category := chi.URLParam(r, "Category")

	result := controller.ArticleService.FindByCategory(r.Context(), Category)
	tmpl := template.Must(template.ParseFiles("./view/category.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
	tmpl.Execute(w, model.DataToRender{Data: result, Path: "/category/" + Category})
}
