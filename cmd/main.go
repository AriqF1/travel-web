package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AriqF1/travel-web/internal/user"
	"github.com/AriqF1/travel-web/internal/auth"
	"github.com/AriqF1/travel-web/internal/vehicle"
	"github.com/AriqF1/travel-web/internal/schedule"
	"github.com/AriqF1/travel-web/internal/booking"

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
		&vehicle.Vehicle{},
		&schedule.Schedule{},
		&booking.Booking{},
	)
	
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("/api")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", auth.RegisterHandler)
		authGroup.POST("/login", auth.LoginHandler)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", auth.ProfileHandler)

		protected.POST("/vehicles", vehicle.CreateVehicleHandler)
		protected.GET("/vehicles", vehicle.GetVehicleHandler)
		protected.GET("/vehicles/:id", vehicle.GetVehicleByIDHandler)
		protected.PUT("/vehicles/:id", vehicle.UpdateVehicleHandler)
		protected.DELETE("/vehicles/:id", vehicle.DeleteVehicleHandler)

		protected.GET("/schedules", schedule.GetScheduleHandler)
		protected.POST("/schedules", schedule.CreateScheduleHandler)
		protected.GET("/schedules/:id", schedule.GetScheduleByIDHandler)
		protected.PUT("/schedules/:id", schedule.UpdateScheduleHandler)
		protected.DELETE("/schedules/:id", schedule.DeleteScheduleHandler)

		protected.POST("/bookings", booking.CreateBookingHandler)
		protected.GET("/bookings", booking.GetBookingsHandler)
		protected.GET("/bookings/:id", booking.GetBookingsByIdHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Erlina Transport API",
		})
	})

	r.Run(":8080")
}