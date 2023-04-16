package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	moviesFile, err := os.Open("movie.csv")
	if err != nil {
		panic(err)
	}

	streamReader := bufio.NewReader(moviesFile)
	streamBuffer := make([]byte, 1000)

	for {
		filePos, err := streamReader.Read(streamBuffer)
		if err != nil {
			break
		}

		// _ = filePos

		// fmt.Println(filePos)
		fmt.Println(string(streamBuffer[:filePos]))
		break
	}
}