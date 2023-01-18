package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func HttpServer() int {
	err := godotenv.Load("/home/asher/repos/Go_Expert/api-tmdb/internal/.env")
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

		movie, err := UnmarshalMovieTmdb(body)
		if err != nil {
			panic(err)
		}

		urlImage := "https://image.tmdb.org/t/p/w500/"

		fmt.Println(urlImage + movie.Results[2].BackdropPath)

	})

	http.ListenAndServe(":8080", nil)
	return http.StatusOK
}
