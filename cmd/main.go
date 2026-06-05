package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AriqF1/travel-web/internal/user"
	"github.com/AriqF1/travel-web/internal/auth"
	"github.com/AriqF1/travel-web/pkg/database"
	"github.com/AriqF1/travel-web/pkg/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	err = database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	database.DB.AutoMigrate(
		&user.User{},
	)

	r := gin.Default()

	api := r.Group("/api")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", auth.RegisterHandler)
		authGroup.POST("/login", auth.LoginHandler)
	}

	protected := api.Group("/profile")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", auth.ProfileHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Erlina Transport API",
		})
	})

	r.Run(":8080")
}