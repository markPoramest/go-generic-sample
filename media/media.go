package media

type Media interface {
	Title() string
	Author() string
	Price() float32
	SetTitle(string)
	SetAuthor(string)
	SetPrice(float32)
	Print()
}
