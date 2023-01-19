package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

<<<<<<< HEAD
var Status *int

func HandleMovies(w http.ResponseWriter, r *http.Request) []byte {
	err := godotenv.Load("internal/.env")
=======
func HttpServer() int {
	err := godotenv.Load("/home/asher/repos/Go_Expert/api-tmdb/internal/.env")
>>>>>>> 9b23fee4d2ef7cbb962c761943954694b3c81c8e
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("API_KEY")

<<<<<<< HEAD
	resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
=======
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
>>>>>>> 9b23fee4d2ef7cbb962c761943954694b3c81c8e

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
