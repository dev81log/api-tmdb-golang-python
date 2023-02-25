package internal

import (
	"log"
)

var result []Result

func InsertResultsIntoTmdb() {

	moviesList := HandleMovies(nil, nil)
	value, err := UnmarshalMovieTmdb([]byte(moviesList))
	if err != nil {
		log.Fatal(err)
	}

	result = append(result, value.Results...)

	for _, r := range result {
		_, err := db.Exec("INSERT INTO results (adult, backdrop_path, id, original_language, original_title, overview, popularity, poster_path, release_date, title, video, vote_average, vote_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", r.Adult, r.BackdropPath, r.ID, r.OriginalLanguage, r.OriginalTitle, r.Overview, r.Popularity, r.PosterPath, r.ReleaseDate, r.Title, r.Video, r.VoteAverage, r.VoteCount)
		if err != nil {
			log.Fatal(err)
		}
	}

}
