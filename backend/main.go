package main

import (
	"net/http"

	"tgtc/backend/database"
	"tgtc/backend/handlers"

	"github.com/gorilla/mux"
)



func main() {
	// init db disini dulu
	database.InitDB()

	// init router
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/banners/{id}", handlers.GetBanner).Methods(http.MethodGet)
	router.HandleFunc("/banners", handlers.MakeBanner).Methods(http.MethodPost)
	router.HandleFunc("/banners/{id}", handlers.UpdateBanner).Methods(http.MethodPatch)

	http.ListenAndServe(":8000", router)
}