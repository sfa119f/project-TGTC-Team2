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
	fmt.Println(idstring)
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

// takes banner info from id in url
func GetBanner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idstr := params["id"]
	fmt.Println(idstr)
	idInt64, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	res, err := service.GetBanner(idInt64)

	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, 
			Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: res, 
			Error: dictionary.NoError,
		})
	}
}

func MakeBanner(w http.ResponseWriter, r *http.Request) {
	banner := dictionary.Banner{}
	json.NewDecoder(r.Body).Decode(&banner)

	err := service.InsertBanner(banner)
	if err != nil {
		fmt.Println("err insert banner:", err)
	}
	
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: banner, Error: dictionary.NoError})
	}
}

func UpdateBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.InvalidParamError})
	}

	banner := dictionary.Banner{Id: id}
	json.NewDecoder(r.Body).Decode(&banner)

	err = service.UpdateBanner(banner)
	if err != nil {
		fmt.Println("err update banner:", err)
	}
	
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: banner, Error: dictionary.NoError})
	}
}
