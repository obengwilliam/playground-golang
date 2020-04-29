package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var dir string
	var defaultPort = "3000"

	port := flag.String("port", defaultPort, "add port to serve file server")
	path := flag.String("path", "", "add directory to server from")
	flag.Parse()

	address := ":" + *port

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, http.FileServer(http.Dir(dir))))

}
