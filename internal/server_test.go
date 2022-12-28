package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	assert.Equal(t, 200, HttpServer())
}
