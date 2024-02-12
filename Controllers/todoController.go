package Controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/Database"
	"go_crud/Database/Models"
	"net/http"
)

func Index(c *gin.Context) {
	var todos []Models.Todo

	if err := Database.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var todo Models.Todo

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully created", "data": todo})
}

func Show(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	if err := Database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	if err := Database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully updated", "data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	if err := Database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	if err := Database.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Todo successfully deleted"})
}
