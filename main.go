package main

import (
	"fmt"

	"example.com/go-crud-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books/:id", controller.ReadBook)
	r.GET("/books", controller.ReadBooks)
	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)
	r.Run(":5000")
}
