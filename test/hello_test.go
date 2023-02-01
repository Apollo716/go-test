package test

import (
	"testing"

	"github.com/Apollo716/go-test/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
