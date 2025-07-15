package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/umer/todo-app/internal/auth"
	"github.com/umer/todo-app/internal/db"
	"github.com/umer/todo-app/internal/todo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Failed to load .env")
	}

	db.InitDB()
	db.DB.AutoMigrate(&auth.User{}, &todo.Todo{})

	log.Println("🚀 Starting server...")

	router := gin.Default()

	auth.RegisterRoutes(router)
	todo.RegisterRoutes(router)

	log.Println("✅ Routes registered. Running server at :8080")
	router.Run(":8080")
}
