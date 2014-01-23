package main

import (
	//"fmt"
	"github.com/rcrowley/go-tigertonic/mocking"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	code, _, response, err := getUser(
		mocking.URL(mux, "GET", "/user/ID"),
		mocking.Header(nil),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
	if code != http.StatusOK {
		t.Fatal(code)
	}
	if response.FirstName != "Bob" {
		t.Errorf("expected Bob, got %v", response.FirstName)
		t.Fatal(response)
	}
}
