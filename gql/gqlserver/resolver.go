package gqlserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tgtc/backend/dictionary"
	"time"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	APIEndpoint string
}

func NewResolver(APIEndpoint string) *Resolver {
	return &Resolver{APIEndpoint: APIEndpoint}
}

func (r *Resolver) GetUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)

		resp, err := http.Get(r.APIEndpoint + "/users/" + fmt.Sprint(id))
		if err != nil {
			fmt.Println("err api call getuser:", err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		userResp := struct {
			Data dictionary.User
			Error string
		}{}
		respJsonString := userResp
		json.Unmarshal(body, &respJsonString)

		return respJsonString.Data, err
	}
}

func (r *Resolver) GetBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["banner_id"].(int)

		res, err := http.Get(r.APIEndpoint + "/banners/" + fmt.Sprint(id))
		if err != nil {
			fmt.Println("err api call GetBanner:", err)
		}
		body, err := ioutil.ReadAll(res.Body)

		bannerRes := struct {
			Data dictionary.Banner
			Error string
		}{}

		respJsonString := bannerRes
		json.Unmarshal(body, &respJsonString)

		return respJsonString.Data, err
	}
}

func (r *Resolver) GetAllBannerOfUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		user_id, _ := p.Args["user_id"].(int)
		resp, err := http.Get(r.APIEndpoint + "/usersXbanners?userId=" + strconv.Itoa(user_id))

		if err != nil {
			fmt.Println("err api call getbannerofuser:", err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		userResp := struct {
			Data []dictionary.Banner
			Error string
		}{}
		respJsonString := userResp
		json.Unmarshal(body, &respJsonString)

		return respJsonString.Data, err
	}
}

func (r *Resolver) CreateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		banner_name, _ := p.Args["banner_name"].(string)
		banner_image, _ := p.Args["banner_image"].(string)
		banner_url, _ := p.Args["banner_url"].(string)
		date_start, _ := p.Args["date_start"].(time.Time)
		date_end, _ := p.Args["date_end"].(time.Time)

		req := dictionary.Banner{
			Name: banner_name,
			Image: banner_image,
			Url: banner_url,
			DateStart: date_start,
			EndDate: date_end,
		}

		newBanner, err := json.Marshal(req)
		if err != nil {
			fmt.Println("err marshalling banner into json str:", err)
		}

		resp, err := http.Post(r.APIEndpoint + "/banners", "application/json", bytes.NewReader(newBanner))
		if err != nil {
			fmt.Println("err api call getuser:", err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		bannerRes := struct {
			Data dictionary.Banner
			Error string
		}{}
		respJsonString := bannerRes
		json.Unmarshal(body, &respJsonString)

		return respJsonString.Data, err
	}
}
