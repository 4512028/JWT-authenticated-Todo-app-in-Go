package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umer/todo-app/internal/db"
)

func CreateTodo(c *gin.Context) {
	userID := c.GetUint("user_id")

	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.UserID = userID

	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func ListTodos(c *gin.Context) {
	userID := c.GetUint("user_id")

	var todos []Todo
	if err := db.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	var todo Todo
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&Todo{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
