package booking

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateBookingHandler(c *gin.Context) {
	var req BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDData, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "user tidak terautentikasi"})
        return
    }

	userID, ok := userIDData.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "tipe data user_id tidak valid"})
        return
    }

	err := CreateBooking(userID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Booking created successfully",
    })
}

	