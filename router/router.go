package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	portString := os.Getenv("PORT")
	fmt.Printf("Server Listening on PORT %s", portString)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Your current IP address is %s", r.RemoteAddr)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})

	return router
}
