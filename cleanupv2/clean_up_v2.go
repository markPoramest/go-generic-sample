package cleanupv2

import (
	"go-generic/media"
	"strings"
)

func CleanUp[T media.Media](i []T) {
	for _, v := range i {
		title := strings.Trim(v.Title(), " ")
		author := strings.Trim(v.Author(), " ")

		title = strings.ToUpper(title)
		author = strings.ToUpper(author)

		v.SetTitle(title)
		v.SetAuthor(author)
	}
}
