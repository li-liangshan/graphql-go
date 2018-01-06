package graphQL

import "github.com/graphql-go/graphql"
import "fmt"

// ExecuteQuery is GraphQL exec entry
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
