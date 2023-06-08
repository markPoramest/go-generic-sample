package media

import "fmt"

type Book struct {
	title  string
	author string
	price  float32
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Price() float32 {
	return b.price
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetPrice(price float32) {
	b.price = price
}

func (b *Book) Print() {
	fmt.Println("Book: " + b.title + " by " + b.author + " with price " + fmt.Sprint(b.price))
}

func NewBook(title string, author string, price float32) *Book {
	return &Book{
		title:  title,
		author: author,
		price:  price,
	}
}
