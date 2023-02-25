DROP TABLE IF EXISTS results;

CREATE TABLE
    results (
        id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
        aldult boolean DEFAULT FALSE NOT NULL,
        backdrop_path VARCHAR(255) NOT NULL,
        original_language VARCHAR(255) NOT NULL,
        original_title VARCHAR(1000) NOT NULL,
        overview VARCHAR(1000) NOT NULL,
        popularity FLOAT NOT NULL,
        poster_path VARCHAR(255) NOT NULL,
        release_date VARCHAR(255) NOT NULL,
        title VARCHAR(255) NOT NULL,
        video boolean DEFAULT FALSE NOT NULL,
        vote_average FLOAT NOT NULL,
        vote_count INT NOT NULL
    );
