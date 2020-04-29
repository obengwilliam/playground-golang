package main

import (
	// "io/ioutil"
	// "log"
	"fmt"
	"net/http"
)

func handleDefaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the default")
}

func main() {

	// resp, err := http.Get("https://kudobuzz.com")

	// if err != nil {
	// 	log.Println("An error occured")
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// log.Println("Respose details")
	// log.Print(string(body))

	http.HandleFunc("/", handleDefaultRoute)
	http.ListenAndServe(":8000", nil)
}
