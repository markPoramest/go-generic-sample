package media

import "fmt"

type Movie struct {
	title  string
	author string
	price  float32
	rating float32
}

func (m *Movie) Title() string {
	return m.title
}

func (m *Movie) Author() string {
	return m.author
}

func (m *Movie) Price() float32 {
	return m.price
}

func (m *Movie) Rating() float32 {
	return m.rating
}

func (m *Movie) SetTitle(title string) {
	m.title = title
}

func (m *Movie) SetAuthor(author string) {
	m.author = author
}

func (m *Movie) SetPrice(price float32) {
	m.price = price
}

func (m *Movie) SetRating(rating float32) {
	m.rating = rating
}

func (m *Movie) Print() {
	fmt.Println("Movie: " + m.title + " by " + m.author + " with price " + fmt.Sprint(m.price) + " and rating " + fmt.Sprint(m.rating))
}

func NewMovie(title string, author string, price float32, rating float32) *Movie {
	return &Movie{
		title:  title,
		author: author,
		price:  price,
		rating: rating,
	}
}
