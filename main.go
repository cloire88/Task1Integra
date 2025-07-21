package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Bed", Completed: false},
	{ID: "2", Item: "Clean Room", Completed: false},
	{ID: "3", Item: "Read Book", Completed: true},
}

func GetToDoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func GetToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{
		"data":    todos,
		"message": "ToDo found",
		"success": true,
	})
}

func AddToDo(context *gin.Context) {
	var newToDo todo

	if err := context.BindJSON(&newToDo); err != nil {
		return
	}

	todos = append(todos, newToDo)

	context.IndentedJSON(http.StatusCreated, newToDo)
}

func GetToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetToDoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "ToDo found", "todo": todo})
}

func toggleToDoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetToDoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", GetToDos)
	router.GET("/todos/:id", GetToDo)
	router.PATCH("/todos/:id", toggleToDoStatus)
	router.POST("/todos", AddToDo)
	router.Run("localhost:9090")
}
