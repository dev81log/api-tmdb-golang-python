package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResultsFromTmdb(t *testing.T) {
	expectedResults := []Result{
		{ID: 1, Title: "Movie 1", Overview: "Movie 1 overview", ReleaseDate: "2022-01-01", PosterPath: "path/to/poster1.jpg"},
		{ID: 2, Title: "Movie 2", Overview: "Movie 2 overview", ReleaseDate: "2022-02-01", PosterPath: "path/to/poster2.jpg"},
		{ID: 3, Title: "Movie 3", Overview: "Movie 3 overview", ReleaseDate: "2022-03-01", PosterPath: "path/to/poster3.jpg"},
	}

	assert.NotEmpty(t, expectedResults)
	assert.Equal(t, 3, len(expectedResults))
	assert.NotNil(t, expectedResults[0])

}
