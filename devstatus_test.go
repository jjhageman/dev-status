package main

import (
	"github.com/rcrowley/go-tigertonic/mocking"
	//"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	code, _, response, err := getUser(
		mocking.URL(mux, "GET", ""),
		mocking.Header(nil),
		nil)
}
