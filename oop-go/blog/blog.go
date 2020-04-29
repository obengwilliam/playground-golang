package blog

import (
	"fmt"
)

type Author struct {
	FirstName string
	LastName  string
	Bio       string
}

type Post struct {
	Author
	Title   string
	Content string
}

func (a Author) FullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

func (p Post) Details() {
	fmt.Println("Title", p.Title)
	fmt.Println("Content", p.Content)
	fmt.Println("Author FullName", p.FullName())
	fmt.Println("Author bio", p.Bio)
}
