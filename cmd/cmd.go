// cmd/server/main.go
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/auth"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Initialize Google OAuth
	auth.InitGoogleOAuth()

	// Initialize database
	config.ConnectDatabase()

	// Routes
	r.GET("/auth/login", auth.HandleGoogleLogin)
	r.GET("/auth/callback", auth.HandleGoogleCallback)

	r.Run(":8080")
}
