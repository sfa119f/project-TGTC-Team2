package gqlserver

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Description: "Detail of a user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"balance": &graphql.Field{
				Type: graphql.Int,
			},
			"member": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BannerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Banner",
		Description: "Detail of a banner",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"dateStart": &graphql.Field{
				Type: graphql.DateTime,
			},
			"dateEnd": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var UserXBannerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserXBanner",
		Description: "Detail of a userXbanner",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"banner_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
