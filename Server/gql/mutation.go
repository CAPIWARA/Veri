package gql

import (
	"Veri/Server/gql/fields"

	"github.com/graphql-go/graphql"
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Mutation",
		Fields: graphql.Fields{
			"createTraject": fields.CreateTrajectory,
		}})
