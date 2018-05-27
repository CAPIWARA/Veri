package gql

import "github.com/graphql-go/graphql"

//Schema is
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})
