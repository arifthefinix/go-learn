package Routes

import (
	"github.com/gin-gonic/gin"
	"go_crud/Controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todo := router.Group("/todo")
	{
		todo.GET("/", Controllers.Index)
		todo.POST("/", Controllers.CreateTodo)
		todo.GET("/:id", Controllers.Show)
		todo.PUT("/:id", Controllers.UpdateTodo)
		todo.DELETE("/:id", Controllers.DeleteTodo)
	}
	return router
}
