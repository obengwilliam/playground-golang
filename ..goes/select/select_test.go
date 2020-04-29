package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {
// 	var slowUrl = "https://facebook.com"
// 	var fastUrl = "http://www.quii.co.uk"

// 	want := fastUrl
// 	got := Racer(slowUrl, fastUrl)

// 	if got != want {
// 		t.Errorf("got %s but want %s", got, want)
// 	}

// }

func makeDelayedServer(waitTime time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(waitTime)
		w.WriteHeader(http.StatusCreated)
	}))
}
func TestRacer(t *testing.T) {
	t.Run("should return fast url", func(t *testing.T) {

		slowServer := makeDelayedServer(0 * time.Millisecond)
		fastServer := makeDelayedServer(30 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		var slowURL = slowServer.URL
		var fastURL = fastServer.URL
		fmt.Println(slowURL, fastURL, "shiit")
		want := slowURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}

	})
}
