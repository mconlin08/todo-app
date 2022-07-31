package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

// Reads json database and unmarshalls the data to a struct
func ReadTodosJSON() Todos {
	data, err := os.Open("../todo-db/db.json")

	if err != nil {
		log.Fatalln(err)
	}

	jsonData, _ := ioutil.ReadAll(data)

	var todos Todos

	json.Unmarshal(jsonData, &todos)

	fmt.Println("data", todos)

	defer data.Close()

	return todos
}

// Gets all todos
func GetTodos(c *gin.Context) {
	fmt.Println("Getting Todos")

	todos := ReadTodosJSON()

	c.IndentedJSON(http.StatusOK, todos)
}

// Gets a todo by its ID
func GetTodoById(c *gin.Context) {

}

// Create a new todo
func CreateTodo(c *gin.Context) {

}

// Updates a todo
func UpdateTodo(c *gin.Context) {

}

// Deletes a todo
func DeleteTodo(c *gin.Context) {

}
