package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jkim7113/centinal/config"
	"github.com/jkim7113/centinal/controller"
	"github.com/jkim7113/centinal/repository"
	"github.com/jkim7113/centinal/service"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	portString := os.Getenv("PORT")
	fmt.Printf("Server Listening on PORT %s \n", portString)
	//DB Configuration
	Db := config.CreateConnection()

	articleRepository := repository.NewArticleRepository(Db)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	userRepository := repository.NewUserRepository(Db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router.Use(middleware.Recoverer)
	router.Group(func(r chi.Router) {
		NewArticleRouter(r, articleController)
	})
	router.Group(func(r chi.Router) {
		NewUserRouter(r, userController)
	})
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	return router
}
