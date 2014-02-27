package main

import (
	"errors"
	"github.com/jjhageman/dev-status/dev"
	"github.com/rcrowley/go-tigertonic"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	mux *tigertonic.TrieServeMux
)

func init() {
	mux = tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/user/{id}", tigertonic.Timed(tigertonic.Marshaled(getUser), "getUser", nil))
}

func main() {
	tigertonic.NewServer(":"+os.Getenv("PORT"), tigertonic.Logged(mux, nil)).ListenAndServe()
}

func getUser(u *url.URL, h http.Header, rq *UserRequest) (int, http.Header, *UserResponse, error) {
	id := u.Query().Get("id")
	strId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}

	dev, err2 := dev.Find(strId)
	if err2 != nil {
		return http.StatusNotFound, nil, nil, tigertonic.NotFound{errors.New("User not found")}
	}

	return http.StatusOK, nil, &UserResponse{dev.FirstName, dev.LastName, dev.GithubID, dev.Status}, nil
}
