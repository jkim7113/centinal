package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Your current IP address is %s", r.RemoteAddr)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})

	return router
}
