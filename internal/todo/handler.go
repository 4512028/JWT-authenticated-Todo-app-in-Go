package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/umer/todo-app/internal/auth"
)

func RegisterRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos", auth.AuthMiddleware())
	todoGroup.POST("/", CreateTodo)
	todoGroup.GET("/", ListTodos)
	todoGroup.PUT("/:id", UpdateTodo)
	todoGroup.DELETE("/:id", DeleteTodo)
}
