package main

import (
	"net/http"
	"os"

	"github.com/jkim7113/centinal/router"
	"github.com/jkim7113/centinal/util"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	routes := router.NewRouter()
	server := http.Server{
		Handler: routes,
		Addr:    ":" + portString,
	}

	err := server.ListenAndServe()
	util.PanicIfError(err)
}
