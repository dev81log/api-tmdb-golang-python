package internal

import "log"

func InsertResultsIntoTmdb() {
	var result []Result

	result = append(result, Result{
		Adult:            false,
		BackdropPath:     "/8uO0gUM8aNqYLs1OsTBQiXu0fEv.jpg",
		ID:               299534,
		OriginalLanguage: "en",
		OriginalTitle:    "Avengers: Endgame",
		Overview:         "After the devastating events of Avengers: Infinity War, the universe is in ruins. With the help of remaining allies, the Avengers assemble once more in order to reverse Thanos' actions and restore balance to the universe.",
		Popularity:       0,
		PosterPath:       "/or06FN3Dka5tukK1e9sl16pB3iy.jpg",
		ReleaseDate:      "2019-04-24",
		Title:            "Avengers: Endgame",
		Video:            false,
		VoteAverage:      8.3,
		VoteCount:        13948,
	})

	// insert results into results table
	for _, r := range result {
		_, err := db.Exec("INSERT INTO results (adult, backdrop_path, id, original_language, original_title, overview, popularity, poster_path, release_date, title, video, vote_average, vote_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", r.Adult, r.BackdropPath, r.ID, r.OriginalLanguage, r.OriginalTitle, r.Overview, r.Popularity, r.PosterPath, r.ReleaseDate, r.Title, r.Video, r.VoteAverage, r.VoteCount)
		if err != nil {
			log.Fatal(err)
		}
	}

}
