package main

import (
	"net/http"
	"time"
)

func mesureResponse(URL string) time.Duration {
	t := time.Now()
	http.Get(URL)
	return time.Since(t)
}

// Racer finds the fastest among two URLs
func Racer(s, f string) (winner string) {
	// slow start time
	sft := mesureResponse(s)
	fft := mesureResponse(f)

	if fft > sft {
		return f
	}
	return s
}
