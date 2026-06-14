package booking

import (
    "errors"
	"net/http"
	"github.com/gin-gonic/gin"

    "gorm.io/gorm"
    "strconv"
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

func GetBookingsHandler(c *gin.Context) {
    bookings, err := GetBookings()

    if err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bookings not found"})
        return
      }

      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    responses := ToBookingResponses(
        bookings,
    )

    c.JSON(http.StatusOK, gin.H{
        "bookings": responses,
    })
}

func GetBookingsByIdHandler(c *gin.Context){
    idParam := c.Param("id")

    id, err := strconv.ParseUint(idParam, 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid Schedule ID",
        })
        return
    }

    booking, err := GetBookingByID(uint(id))

    if err != nil {

        if errors.Is(err, gorm.ErrRecordNotFound){
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Booking not found",
            })
            return
        }
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
        return
    }   


    response := ToBookingResponse(booking)
    c.JSON(http.StatusOK, gin.H{
        "booking" : response,
    })
}
	