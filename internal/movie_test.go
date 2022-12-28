package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMovieTmdbError(t *testing.T) {
	data := []byte(`{invalid json`)
	_, err := UnmarshalMovieTmdb(data)
	assert.Error(t, err)
}

func TestUnmarshalMovieTmdb(t *testing.T) {
	data := []byte(`{
		"page": 1,
		"results": [
			{
				"adult": false,
				"backdrop_path": "/path/to/backdrop.jpg",
				"genre_ids": [28, 12, 16],
				"id": 12345,
				"original_language": "en",
				"original_title": "Movie Title",
				"overview": "Movie overview",
				"popularity": 123.456,
				"poster_path": "/path/to/poster.jpg",
				"release_date": "2022-01-01",
				"title": "Movie Title",
				"video": false,
				"vote_average": 7.8,
				"vote_count": 1000
			}
		],
		"total_pages": 10,
		"total_results": 100
	}`)
	expected := MovieTmdb{
		Page:         1,
		TotalPages:   10,
		TotalResults: 100,
		Results: []Result{
			{
				Adult:            false,
				BackdropPath:     "/path/to/backdrop.jpg",
				GenreIDS:         []int64{28, 12, 16},
				ID:               12345,
				OriginalLanguage: "en",
				OriginalTitle:    "Movie Title",
				Overview:         "Movie overview",
				Popularity:       123.456,
				PosterPath:       "/path/to/poster.jpg",
				ReleaseDate:      "2022-01-01",
				Title:            "Movie Title",
				Video:            false,
				VoteAverage:      7.8,
				VoteCount:        1000,
			},
		},
	}

	result, err := UnmarshalMovieTmdb(data)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(result, expected))
}
