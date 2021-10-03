package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type Resolver struct {}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// id, _ := p.Args["product_id"].(int)

		// update to use Usecase from previous session
		return nil, nil
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
