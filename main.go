package main

import (
	"net/http"

	"github.com/dev81log/api-tmdb/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/movies", routerMovies)

	http.ListenAndServe(":8080", r)
}

func routerMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(internal.HandleMovies(w, r)))

}
