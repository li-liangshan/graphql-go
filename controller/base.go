package controller

import (
	"encoding/json"
	"fmt"
	"go-graphql-api/graphQL"
	"net/http"
)

func GraphQLController(w http.ResponseWriter, r *http.Request) {
	result := graphQL.ExecuteQuery(r.URL.Query().Get("query"), graphQL.Schema)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}
