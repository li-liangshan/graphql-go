package graphQL

import (
	"github.com/graphql-go/graphql"
)

// Schema is graphQL schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
