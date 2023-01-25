package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResultsFromTmdb(t *testing.T) {
	// Set up test data
	expectedResults := []Result{
		{ID: 1, Title: "Movie 1", Overview: "Movie 1 overview", ReleaseDate: "2022-01-01", PosterPath: "path/to/poster1.jpg"},
		{ID: 2, Title: "Movie 2", Overview: "Movie 2 overview", ReleaseDate: "2022-02-01", PosterPath: "path/to/poster2.jpg"},
		{ID: 3, Title: "Movie 3", Overview: "Movie 3 overview", ReleaseDate: "2022-03-01", PosterPath: "path/to/poster3.jpg"},
	}

	// Insert test data into the database
	for _, result := range expectedResults {
		_, err := db.Exec("INSERT INTO results (id, title, overview, release_date, poster_path) VALUES (?, ?, ?, ?, ?)", result.ID, result.Title, result.Overview, result.ReleaseDate, result.PosterPath)
		assert.Nil(t, err)
	}

	// Call the function to test
	actualResults := GetResultsFromTmdb()

	// Compare the expected and actual results
	assert.Len(t, actualResults, len(expectedResults))
	for i := range actualResults {
		assert.Equal(t, actualResults[i], expectedResults[i])
	}

	// Clean up the test data
	_, err := db.Exec("DELETE FROM results")
	assert.Nil(t, err)
}
