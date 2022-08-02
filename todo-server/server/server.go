package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Order       int    `json:"order"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

var todosList Todos

// Sets todoList variable in the server
func InitializeDB() {
	todosList = ReadTodosJSON()
}

// Reads json database and unmarshalls the data to a struct
func ReadTodosJSON() Todos {
	fmt.Println("Reading JSON DB...")
	data, err := os.Open("../todo-db/db.json")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	jsonData, jsonErr := ioutil.ReadAll(data)
	if jsonErr != nil {
		log.Fatalln("Error reading json file:", jsonErr)
	}

	var todos Todos

	json.Unmarshal(jsonData, &todos)

	fmt.Println("data", todos)

	defer data.Close()

	return todos
}

// Saves todo data to the json db
func SaveData(todo *Todo) {
	fmt.Println("\nSaving...")

	newTodoList := append(todosList.Todos, *todo)

	todosList.Todos = newTodoList

	file, _ := json.MarshalIndent(todosList, "", " ")

	_ = ioutil.WriteFile("../todo-db/db.json", file, 0644)

	fmt.Println("Save Complete")
}

// Updates todo data in the json db
func UpdateData() {
	fmt.Println("\nUpdating...")

	file, _ := json.MarshalIndent(todosList, "", " ")

	_ = ioutil.WriteFile("../todo-db/db.json", file, 0644)

	fmt.Println("Update Complete")
}

// Gets all todos
func GetAllTodos(c *gin.Context) {
	fmt.Println("Getting Todos")

	c.IndentedJSON(http.StatusOK, todosList)
}

// Gets a todo by its ID
func GetTodoById(c *gin.Context) {
	id := c.Query("id")
	var todo Todo

	for _, t := range(todosList.Todos) {
		if t.Id == id {
			todo = t
		}
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// Creates a new todo
func CreateTodo(c *gin.Context) {
	id := c.DefaultQuery("id", cuid.New())
	completed := c.Query("completed")
	var isTodoCompleted bool

	if completed == "true" {
		isTodoCompleted = true
	} else {
		isTodoCompleted = false
	}

	todo := &Todo{
		Id:          id,
		Title:       c.Query("title"),
		Description: c.Query("description"),
		Completed:   isTodoCompleted,
		Order: len(todosList.Todos) + 1,
	}

	fmt.Println("Todo Created!")

	SaveData(todo)

	c.IndentedJSON(http.StatusCreated, todo)
}

// Updates a todo
func UpdateTodo(c *gin.Context) {
	id := c.Query("id")
	completed := c.Query("completed")
	var isTodoCompleted bool

	if completed == "true" {
		isTodoCompleted = true
	} else {
		isTodoCompleted = false
	}

	todo := &Todo{
		Id:          id,
		Title:       c.Query("title"),
		Description: c.Query("description"),
		Completed:   isTodoCompleted,
		Order: 		 0,
	}

	for i, t := range(todosList.Todos) {
		if t.Id == todo.Id {
			todo.Order = i + 1
			todosList.Todos[i] = *todo
		}
	}
	
	UpdateData()

	c.IndentedJSON(http.StatusOK, todo)
}

func RemoveTodo(slice []Todo, index int) []Todo {
	return append(slice[:index], slice[index+1:]...)
}

// Deletes a todo
func DeleteTodo(c *gin.Context) {
	id := c.Query("id")

	for i, t := range(todosList.Todos) {
		if id == t.Id {
			todosList.Todos = RemoveTodo(todosList.Todos, i)
		}
	}

	UpdateData()

	c.IndentedJSON(http.StatusOK, id)
}
