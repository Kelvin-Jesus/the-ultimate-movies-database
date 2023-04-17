package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	moviesFile, err := os.Open("movie.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(moviesFile)

	for {
		csvLine, err := csvReader.Read()
		
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		movieId := csvLine[:1]
		movieName := csvLine[1:2]
		movieGenres := csvLine[2:]
		fmt.Println(movieId, movieGenres, movieName)
		fmt.Println(strings.Replace(queryInsertMovie(), "{MOVIE_ID}", movieId[0], -1))
		fmt.Println(strings.Replace(queryInsertMovie(), "{MOVIE_NAME}", movieName[0], -1))
		fmt.Println(strings.Replace(queryInsertMovie(), "{MOVIE_GENRES}", strings.Join(movieGenres, ","), -1))
		// break
	}
}

func queryInsertMovie() string {
	return `INSERT INTO movies (movie_id, movie_name, movie_genres) VALUES ({MOVIE_ID}, {MOVIE_NAME}, {MOVIE_GENRES})`
}