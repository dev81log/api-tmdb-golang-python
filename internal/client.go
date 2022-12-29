package internal

import (
	"net/http"
)

func ClientMovie() (*http.Request, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/movie", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resp.

	return req, err


}


