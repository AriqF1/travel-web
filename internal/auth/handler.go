package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := Register(req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func LoginHandler(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}	

	token, err := Login(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}	
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})	
}

func ProfileHandler(c *gin.Context) {

	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	c.JSON(200, gin.H{
		"user_id": userID,
		"email": email,
	})
}