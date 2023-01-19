package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	status := HandleMovies
	assert.NotNil(t, status, "Expected status to be not nil")
	assert.NotEmpty(t, status, "Expected status to be not empty")
}
