package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func (s *ServerPool) HealthCheck() {
	for _, b := range s.backends {
		log.Println(b)
	}
}

func ld(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var serverList string
	var port int

	flag.StringVar(&serverList, "backends", "", "Load balanced backends, use commas to seperate")
	flag.IntVar(&port, "port", 3030, "Port to server")

	flag.Parse()

	log.Println("serverlist:", serverList)
	log.Println("port:", port)

	if len(serverList) == 0 {
		log.Fatal("Please provide backend servers")
	}

	server := http.Server{Addr: fmt.Sprintf(":%d", port), Handler: http.HandlerFunc(ld)}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
