package internal

import (
	"net/http"
)

func HttpServer() int {
	http.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=2e93e66834b37efaaf60641340a1f6eb")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	})

	http.ListenAndServe(":8080", nil)
	return http.StatusOK
}
