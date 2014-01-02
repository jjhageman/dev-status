package main

import (
	"./dev"
	"encoding/json"
	"net/http"
)

//var devs = dev.NewDevManager()

const PathPrefix = "/dev/"

func allDevsHandler(w http.ResponseWriter, r *http.Request) {
	res := struct{ Devs []*dev.Dev }{}
	return json.NewEncoder(w).Encode(res)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc(PathPrefix+"{id}", allDevsHandler)
	http.ListenAndServe(":8080", nil)
}
