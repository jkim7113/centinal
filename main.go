package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	router := chi.NewRouter()
	server := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Server Listening on PORT %s", portString)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
