package main

import (
	"github.com/jjhageman/dev-status/db"
	"github.com/jjhageman/dev-status/dev"
	"github.com/rcrowley/go-tigertonic"
	"net/http"
	"net/url"
)

var (
	mux *tigertonic.TrieServeMux
)

func init() {
	dev.Dbmap = db.InitDb()
	mux = tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/user/{id}", tigertonic.Timed(tigertonic.Marshaled(getUser), "getUser", nil))
}

func main() {
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}

func getUser(u *url.URL, h http.Header, rq *UserRequest) (int, http.Header, *UserResponse, error) {
	return http.StatusOK, nil, &UserResponse{"Test", "User", "1234", "Available"}, nil
}
