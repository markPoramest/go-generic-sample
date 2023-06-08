package cleanupv1

import (
	"go-generic/media"
	"strings"
)

func CleanUpBooks(b []*media.Book) []media.Book {
	var cleanedBook []media.Book

	for _, v := range b {
		title := strings.Trim(v.Title(), " ")
		author := strings.Trim(v.Author(), " ")

		title = strings.ToUpper(title)
		author = strings.ToUpper(author)

		book := media.NewBook(title, author, v.Price())
		cleanedBook = append(cleanedBook, *book)
	}

	return cleanedBook
}

func CleanUpMovies(m []*media.Movie) []media.Movie {
	var cleanedMovie []media.Movie

	for _, v := range m {
		title := strings.Trim(v.Title(), " ")
		name := strings.Trim(v.Author(), " ")

		movie := media.NewMovie(title, name, v.Price(), v.Rating())
		cleanedMovie = append(cleanedMovie, *movie)
	}

	return cleanedMovie
}
