package main

import dp "github.com/obengwilliam/go-playground/dependencyinjection"
import "net/http"

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dp.Greet(w, "worlds")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
