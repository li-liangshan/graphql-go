package router

import (
	"go-graphql-api/controller"
	"go-graphql-api/graphQL"
	"net/http"

	"github.com/graphql-go/handler"
)

func Run() {
	http.HandleFunc("/graph", controller.GraphQLController)
	hand := handler.New(&handler.Config{
		Schema:   &graphQL.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	// serve HTTP
	http.Handle("/graphql", hand)
	http.ListenAndServe(":8080", nil)
}
