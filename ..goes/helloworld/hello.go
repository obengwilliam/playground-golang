package main

import (
	"fmt"
)

const englishHelloPrefix = "hello "
const frenchHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Holla "

func getPrefix(language string) string {
	var prefix = englishHelloPrefix
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	}
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}
func main() {
	fmt.Println(Hello("Chris", "en"))
}
