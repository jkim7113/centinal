package main

import (
	"net/http"
	"os"

	"github.com/jkim7113/centinal/config"
	"github.com/jkim7113/centinal/controller"
	"github.com/jkim7113/centinal/repository"
	"github.com/jkim7113/centinal/router"
	"github.com/jkim7113/centinal/service"
	"github.com/jkim7113/centinal/util"
	"github.com/joho/godotenv"
)

func main() {
	//Load environment variables from .env
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	//DB Configuration
	Db := config.CreateConnection()

	articleRepository := repository.NewArticleRepository(Db)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	routes := router.NewRouter(articleController)
	server := http.Server{
		Handler: routes,
		Addr:    ":" + portString,
	}

	err := server.ListenAndServe()
	util.PanicIfError(err)
}
