package fields

import (
	"Veri/Server/gql/resolvers"
	"Veri/Server/gql/types"

	"github.com/graphql-go/graphql"
)

//ContractQuery is
var ContractQuery = &graphql.Field{
	Type:        types.Contract,
	Description: "get contract by text",
	Args: graphql.FieldConfigArgument{
		"text": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.GetContract,
}

//CreateTrajectory is
var CreateTrajectory = &graphql.Field{
	Type:        types.Contract,
	Description: "create contract",
	Args: graphql.FieldConfigArgument{
		"text": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.CreateTrajectory,
}
