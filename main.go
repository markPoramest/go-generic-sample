package main

import (
	"fmt"
	"go-generic/cleanupv1"
	"go-generic/cleanupv2"
	"go-generic/media"
)

func main() {
	books := []*media.Book{
		media.NewBook(" the hitchhiker's guide to the galaxy        ", "Douglas Adams", 9.21),
		media.NewBook("   The Hobbit", "J.R.R. Tolkien", 7.53),
	}

	movies := []*media.Movie{
		media.NewMovie("      Back to the Future   ", "Robert Zemeckis", 7.42, 8.5),
		media.NewMovie("The Godfather      ", "Francis Ford Coppola", 8.53, 9.2),
	}

	cleanedMovies := cleanupv1.CleanUpMovies(movies)
	cleanedBooks := cleanupv1.CleanUpBooks(books)

	fmt.Println("========= Cleaned v1 =========")

	for _, v := range cleanedBooks {
		v.Print()
	}

	for _, v := range cleanedMovies {
		v.Print()
	}

	cleanupv2.CleanUp(movies)
	cleanupv2.CleanUp(books)

	fmt.Println("========= Cleaned v2 =========")

	for _, v := range books {
		v.Print()
	}

	for _, v := range movies {
		v.Print()
	}

}
