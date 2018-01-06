package main

import (
	"fmt"
	"go-graphql-api/models"
	"go-graphql-api/router"
)

func main() {
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single todo: curl -g 'http://localhost:8080/graphql?query={todo(id:\"b\"){id,text,done}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
	fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")
	models.Init()
	router.Run()
}


