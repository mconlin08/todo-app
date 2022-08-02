package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	server "todo/todo-server/server"
)

func main() {
	fmt.Println("Starting Server")
	server.InitializeDB()

	router := gin.Default()
	router.GET("/todos", server.GetAllTodos)
	router.GET("/todo", server.GetTodoById)
	router.POST("/todo/create", server.CreateTodo)
	router.PUT("/todo/update", server.UpdateTodo)
	router.DELETE("todo/delete", server.DeleteTodo)
	router.Run("localhost:8080")
}
