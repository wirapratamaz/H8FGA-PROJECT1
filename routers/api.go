package routers

import (
	"books/controllers"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Book API
// @version 1.0
// @description This is a simple services for managing book
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email wirabagus185@gmail.com
// @license.name Apace 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
