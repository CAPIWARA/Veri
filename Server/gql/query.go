package gql

import (
	"hack-bike/gql/fields"

	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "rootQuery",
		Description: "rootQuery",
		Fields: graphql.Fields{
			"contract": fields.ContractQuery,
		},
	},
)
