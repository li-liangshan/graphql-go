package graphQL

import (
	"fmt"
	"go-graphql-api/models"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootQuery",
	Fields: graphql.Fields{
		"todo": &graphql.Field{
			Type:        todoType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := params.Args["id"].(string)

				if isOK {
					for _, todo := range models.TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}
				return models.Todo{}, nil
			},
		},
		"lastTodo": &graphql.Field{
			Type:        todoType,
			Description: "Last todo added",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return models.TodoList[len(models.TodoList)-1], nil
			},
		},

		"todoList": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "List of todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				fmt.Println(models.TodoList)
				return models.TodoList, nil
			},
		},
	},
})
