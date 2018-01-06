package graphQL

import (
	"go-graphql-api/lib"
	"go-graphql-api/models"

	"github.com/graphql-go/graphql"
)

// rootMutation is graphql mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootMutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type:        todoType,
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				text, _ := params.Args["text"].(string)

				newID := lib.RandStringRunes(8)
				newTodo := models.Todo{
					ID:   newID,
					Text: text,
					Done: false,
				}

				models.TodoList = append(models.TodoList, newTodo)
				return newTodo, nil
			},
		},

		"updateTodo": &graphql.Field{
			Type:        todoType,
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				done, _ := params.Args["done"].(bool)
				id, _ := params.Args["id"].(string)
				affectedTodo := models.Todo{}

				for i := 0; i < len(models.TodoList); i++ {
					if models.TodoList[i].ID == id {
						models.TodoList[i].Done = done
						affectedTodo = models.TodoList[i]
						break
					}
				}
				return affectedTodo, nil
			},
		},
	},
})
