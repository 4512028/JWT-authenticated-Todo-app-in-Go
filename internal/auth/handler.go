package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")

	log.Println("🔧 Registering /auth routes")

	authGroup.POST("/register", registerHandler)
	authGroup.POST("/login", loginHandler)

	// ✅ Add this line
	authGroup.GET("/me", AuthMiddleware(), meHandler)

	log.Println("✅ Registered: /auth/register, /auth/login, /auth/me")
}
