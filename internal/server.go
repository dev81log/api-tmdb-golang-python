package internal

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func HttpServer() int {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("API_KEY")

	http.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=" + apiKey)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	})

	http.ListenAndServe(":8080", nil)
	return http.StatusOK
}
