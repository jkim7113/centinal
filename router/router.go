package router

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/controller"
	"github.com/jkim7113/centinal/model"
)

func NewRouter(articleController *controller.ArticleController) *chi.Mux {
	router := chi.NewRouter()
	portString := os.Getenv("PORT")
	fmt.Printf("Server Listening on PORT %s \n", portString)

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
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	return router
}
