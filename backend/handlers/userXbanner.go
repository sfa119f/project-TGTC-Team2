package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tgtc/backend/dictionary"
	"tgtc/backend/service"
)

func MakeUserXBanner(w http.ResponseWriter, r *http.Request) {
	var userXBanner dictionary.User_X_Banner
	json.NewDecoder(r.Body).Decode(&userXBanner)
	
	err := service.InsertUserXBanner(userXBanner)
	if err != nil {
		fmt.Println("err insert user:", err)
	}

	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: userXBanner, Error: dictionary.NoError})
	}
}

func GetBannersOfUser(w http.ResponseWriter, r *http.Request) {
	userIdstring := r.FormValue("userId")
	userIdInt64, err := strconv.ParseInt(userIdstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	banners_id, err := service.GetBannersOfUser(userIdInt64)
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	}

	banners := []dictionary.Banner{}
	for _, banner_id := range banners_id {
		banner, err := service.GetBanner(int64(banner_id))
		if err != nil {
			json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.NotFoundError})
		}
		banners = append(banners, banner)
	}
	
	json.NewEncoder(w).Encode(dictionary.APIResponse{Data: banners, Error: dictionary.NoError})
}
