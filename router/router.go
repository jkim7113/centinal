package router

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/controller"
)

func NewRouter(articleController *controller.ArticleController) *chi.Mux {
	router := chi.NewRouter()
	portString := os.Getenv("PORT")
	fmt.Printf("Server Listening on PORT %s \n", portString)

	router.Get("/", articleController.FindAll)
	router.Get("/category/{Category}", articleController.FindByCategory)
	router.Get("/article/{UUID}", articleController.FindById)
	router.Post("/article", articleController.Create)
	router.Put("/article/{UUID}", articleController.Update)
	router.Get("/new/article", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./view/new_article.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
		tmpl.Execute(w, nil)
	})
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	return router
}
