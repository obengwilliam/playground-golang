package main

import "fmt"

func init() {
	fmt.Println("init in sandbox.go")
}

var _ int64 = s()
var c = 3

func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}

func main() {
	fmt.Println("main")
}
