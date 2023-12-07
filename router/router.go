package router

import (
	"fmt"
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
	router.Post("/articles", articleController.Create)
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	return router
}
