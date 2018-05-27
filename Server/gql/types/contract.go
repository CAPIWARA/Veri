package types

import "github.com/graphql-go/graphql"

//Contract is
var Contract = graphql.NewObject(graphql.ObjectConfig{
	Name: "contract",
	Fields: graphql.Fields{
		"text": &graphql.Field{
			Type:        graphql.String,
			Description: "contract text",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "contract email",
		},
	}})
