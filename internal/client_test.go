package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientMovie(t *testing.T) {
	req, err := ClientMovie()
	assert.NoError(t, err, "ClientMovie returned an error")

	assert.Equal(t, "GET", req.Method, "Expected GET request")
	assert.Equal(t, "http://localhost:8080/movies", req.URL.String(), "Expected request to be sent to http://localhost:8080/movies")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "Error executing request")
	assert.NotEmpty(t, resp, "Expected response to be empty")
	assert.Equal(t, 200, resp.StatusCode, "Expected status code to be 200")

}
