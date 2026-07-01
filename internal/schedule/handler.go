package schedule

import (
	"net/http"
	"strconv"
	"errors"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateScheduleHandler(c *gin.Context) {
	var req ScheduleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := CreateSchedule(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})	

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Schedule created successfully",
	})
}

func GetScheduleHandler(c *gin.Context) {	
	schedules, err := GetSchedules()

	 if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Schedule not found",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	responses := ToScheduleResponses(
		schedules,
	)

	c.JSON(http.StatusOK, gin.H{
		"schedules": responses,
	})
}

func GetScheduleByIDHandler(c *gin.Context) {
    idParam := c.Param("id")

    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid schedule ID",
        })
        return
    }

    schedule, err := GetScheduleByID(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Schedule not found",
            })
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    availableSeats, err := GetAvailableSeats(schedule.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    response := ToScheduleResponse(schedule, availableSeats)
    c.JSON(http.StatusOK, gin.H{
        "schedule": response,
    })
}

func UpdateScheduleHandler(c *gin.Context) {
	var req ScheduleRequest

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid schedule ID",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := UpdateSchedule(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Schedule updated successfully",
	})
}

func DeleteScheduleHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid schedule ID",
		})
		return
	}

	if err := DeleteSchedule(uint(id)); err != nil { 
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Schedule not found",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Schedule deleted successfully",
	})
}

