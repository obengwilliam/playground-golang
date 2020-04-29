package mock

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := bytes.Buffer{}
	want := "3"
	got := CountDown(buffer)

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
