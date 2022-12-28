package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	code := HttpServer()
	assert.Equal(t, 200, code)
}
