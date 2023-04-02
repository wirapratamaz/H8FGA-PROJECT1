package routers

import (
	"books/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Read by Id
	router.GET("/books/:id", controllers.GetOneBook)
	// Read All
	router.GET("/books", controllers.GetAllBooks)
	// Create
	router.POST("/books", controllers.CreateBook)
	// Update by Id
	router.PATCH("/books/:id", controllers.UpdateBook)
	// Delete by Id
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
