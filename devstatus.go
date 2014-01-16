package main

import (
	//"./dev"
	"github.com/rcrowley/go-tigertonic"
	"net/http"
	"net/url"
)

func main() {
	mux := tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/user/{id}", tigertonic.Timed(tigertonic.Marshaled(getUser), "getUser", nil))
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}

func getUser(u *url.URL, h http.Header, rq *UserRequest) (int, http.Header, *UserResponse, error) {
	return http.StatusOK, nil, &UserResponse{}, nil
}
