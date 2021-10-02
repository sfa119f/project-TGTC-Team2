package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("webstatic"))
	r.PathPrefix("/").Handler(fs)

	http.ListenAndServe(":9000", r)
}