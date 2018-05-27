package types

import "github.com/graphql-go/graphql"

//Contract is
var Contract = graphql.NewObject(graphql.ObjectConfig{
	Name: "contract",
	Fields: graphql.Fields{
		"uuid": &graphql.Field{
			Type:        graphql.String,
			Description: "contract text",
		},
		"balance": &graphql.Field{
			Type:        graphql.String,
			Description: "contract balance",
		},
		"trajectory": &graphql.Field{
			Type:        graphql.String,
			Description: "contract trajectory",
		},
	}})
