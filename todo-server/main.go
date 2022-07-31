package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	server "todo/todo-server/server"
)

func main() {
	fmt.Println("Starting Server")

	router := gin.Default()
	router.GET("/todos", server.GetTodos)
	// router.GET("/books", getBooks)
	// router.GET("books/:id", bookById)
	// router.POST("/books", createBook)
	// router.PATCH("/checkout", checkoutBook)
	// router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
