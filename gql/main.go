package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("gql/webstatic"))
	r.PathPrefix("/").Handler(fs)

}