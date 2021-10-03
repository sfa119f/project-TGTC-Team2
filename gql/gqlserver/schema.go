package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "UsersAndBanners",
			Description: "All query related to getting users and banners data",
			Fields: graphql.Fields{
				"UserDetail": &graphql.Field{
					Type:        UserType,
					Description: "Get user by ID",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetUser(),
				},
				"BannerDetail": &graphql.Field{
					Type:        BannerType,
					Description: "Get banner by ID",
					Args: graphql.FieldConfigArgument{
						"banner_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetBanner(),
				},
				"UserBanners": &graphql.Field{
					Type:        graphql.NewList(BannerType),
					Description: "Get banners of a user by userId",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetAllBannerOfUser(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "updateData",
			Description: "mutations related to updating existing data",
			Fields: graphql.Fields{
				"createBanner": &graphql.Field{
					Type: BannerType,
					Description: "Create new banner",
					Args: graphql.FieldConfigArgument{
						"banner_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_image": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_url": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"date_start": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"date_end": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.CreateBanner(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
