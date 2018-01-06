package models

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var TodoList []Todo

func init() {
	todo1 := Todo{ID: "a", Text: "A todo not to forget", Done: false}
	todo2 := Todo{ID: "b", Text: "This is the most important", Done: false}
	todo3 := Todo{ID: "c", Text: "Please do this or else", Done: false}
	TodoList = append(TodoList, todo1, todo2, todo3)
}
