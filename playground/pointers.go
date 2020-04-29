package main

import . "fmt"

func toOne(one int) {
	one = 1
}

func toOneP(one *int) {
	*one = 0
}

func main() {
	one := 2
	toOne(one)
	Println(one)
	toOneP(&one)
	Println(one)
}
