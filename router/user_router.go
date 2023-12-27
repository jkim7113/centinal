package router

import (
	// "html/template"
	// "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jkim7113/centinal/controller"
	// "github.com/jkim7113/centinal/model"
)

func NewUserRouter(router chi.Router, userController *controller.UserController) {
	router.Post("/user", userController.Create)
	router.Route("/user/{UUID}", func(r chi.Router) {
		r.Get("/", userController.FindById)
		r.Put("/", userController.Update)
		r.Delete("/", userController.Delete)
		// r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
		// 	UUID := chi.URLParam(r, "UUID")
		// 	tmpl := template.Must(template.ParseFiles("./view/edit_article.html", "./view/config.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
		// 	tmpl.Execute(w, model.DataToRender{Data: nil, Path: "/article/" + UUID})
		// })
	})
	// router.Route("/session", func(r chi.Router) {
	// 	r.Post("/",)
	// })
}
