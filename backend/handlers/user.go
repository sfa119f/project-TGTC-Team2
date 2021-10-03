package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tgtc/backend/dictionary"
	"tgtc/backend/service"

	"github.com/gorilla/mux"
)

// takes user info from id in url
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idstring := params["id"]
	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	res, err := service.GetUser(idInt64)
	
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: res, Error: dictionary.NoError})
	}
}

func MakeUser(w http.ResponseWriter, r *http.Request) {
	user := dictionary.User{}
	json.NewDecoder(r.Body).Decode(&user)

	err := service.InsertUser(user)
	if err != nil {
		fmt.Println("err insert user:", err)
	}

	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil,
			Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: user,
			Error: dictionary.NoError,
		})
	}
}