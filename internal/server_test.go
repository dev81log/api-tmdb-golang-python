package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
<<<<<<< HEAD
	status := HandleMovies
	assert.NotNil(t, status, "Expected status to be not nil")
	assert.NotEmpty(t, status, "Expected status to be not empty")
=======
	status := HttpServer()
	assert.Equal(t, 200, status)
>>>>>>> 9b23fee4d2ef7cbb962c761943954694b3c81c8e
}
