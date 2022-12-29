package internal

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var BodyJson []byte

func HttpServer() int {
	err := godotenv.Load("/home/asher/repos/Estudos/api-movie-data/api-tmdb/internal/.env")
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

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		UnmarshalMovieTmdb(body)

	})

	http.ListenAndServe(":8080", nil)
	return http.StatusOK
}
