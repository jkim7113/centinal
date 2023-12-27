package router

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/controller"
	"github.com/jkim7113/centinal/model"
)

func NewArticleRouter(router chi.Router, articleController *controller.ArticleController) {
	router.Get("/", articleController.FindAll)
	router.Get("/category/{Category}", articleController.FindByCategory)
	router.Post("/article", articleController.Create)
	router.Route("/article/{UUID}", func(r chi.Router) {
		r.Get("/", articleController.FindById)
		r.Put("/", articleController.Update)
		r.Delete("/", articleController.Delete)
		r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			UUID := chi.URLParam(r, "UUID")
			tmpl := template.Must(template.ParseFiles("./view/edit_article.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
			tmpl.Execute(w, model.DataToRender{Data: nil, Path: "/article/" + UUID})
		})
	})

	router.Get("/new/article", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./view/new_article.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
		tmpl.Execute(w, nil)
	})
}
