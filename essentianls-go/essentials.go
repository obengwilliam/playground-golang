package main

import (
	"fmt"
	"github.com/obengwilliam/essentials-go/socialmedia"
)

func main() {
	keywords := []string{"jump", "good"}

	post := socialmedia.NewPost("william", "https://image.url.com", "https://thumnail.url.image.com", "cool", "https://unique.url.com", "messageBody", socialmedia.Moods["happy"], keywords, []string{"good", "wiliam"})

	fmt.Printf("Mypost %v\n", post)
}
