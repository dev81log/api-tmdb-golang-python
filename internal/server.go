package internal

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var Status *int

func HandleMovies(w http.ResponseWriter, r *http.Request) []byte {
	err := godotenv.Load("internal/.env")
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("API_KEY")

	resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
