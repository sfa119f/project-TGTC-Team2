package gqlserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tgtc/backend/dictionary"

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

func (r *Resolver) GetAllProducts() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	}
}

// func (r *Resolver) UpdateProduct() graphql.FieldResolveFn {
// 	return func(p graphql.ResolveParams) (interface{}, error) {
// 		id, _ := p.Args["product_id"].(int)

// 		// update to use Usecase from previous session
// 		return service.UpdateProduct(id)
// 	}
// }
