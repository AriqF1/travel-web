package vehicle

import (
	"net/http"
	"errors"
	"gorm.io/gorm"
	
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateVehicleHandler(c *gin.Context) {
	var req CreateVehicleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := createVehicle(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Vehicle created successfully",
	})
}

func GetVehicleHandler(c *gin.Context) {
	vehicles, err := GetVehicle()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := ToVehicleResponses(
		vehicles,
	)

	c.JSON(http.StatusOK, gin.H{
		"vehicles": responses,
	})
}

func GetVehicleByIDHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vehicle ID",
		})
		return
	}

	vehicle, err := GetVehicleByID(uint(id))

	if err != nil {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vehicle not found",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}

	c.JSON(http.StatusOK, gin.H{
		"vehicle": vehicle,
	})
}

func DeleteVehicleHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vehicle ID",
		})
		return
	}

	if err := DeleteVehicle(uint(id)); err != nil {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vehicle not found",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vehicle deleted successfully",
	})

}

func UpdateVehicleHandler(c *gin.Context) {
	var req UpdateVehicleRequest

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vehicle ID",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := UpdateVehicle(uint(id), req); err != nil {

    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "vehicle not found",
        })
        return
    }

    c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
    })
    	return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vehicle updated successfully",
	})
}