package main

import (
	"log"
	"net/http"

	"tgtc/gql/gqlserver"

	"github.com/gorilla/mux"
)

const (
	APIEndpoint = "http://localhost:8000"
)

func main() {
	router := mux.NewRouter()

	pResolver := gqlserver.NewResolver(APIEndpoint)
	schemaWrapper := gqlserver.NewSchemaWrapper().
		WithProductResolver(pResolver)

	if err := schemaWrapper.Init(); err != nil {
		log.Fatal("unable to parse schema, err: ", err.Error())
	}

	// GraphQL Handler
	router.Path("/graphql").Handler(gqlserver.NewHandler(schemaWrapper).Handle())

	// this assumes main.go is called from root project,
	// change this accordingly, if it's called elsewhere
	fs := http.FileServer(http.Dir("webstatic"))
	router.PathPrefix("/").Handler(fs)

	http.ListenAndServe(":9000", router)
}
