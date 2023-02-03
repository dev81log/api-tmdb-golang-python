package internal

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	status := HandleMovies
	assert.NotNil(t, status, "Expected status to be not nil")
	assert.NotEmpty(t, status, "Expected status to be not empty")
}

func TestHandleMovies(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleMovies(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}
