package main

import (
	"net/http"

	"tgtc/backend/handlers"

	"github.com/gorilla/mux"
)



func main() {
	// init db disini dulu
	// database.Init()

	// init router
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/banners/{id}", handlers.GetBanner).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}