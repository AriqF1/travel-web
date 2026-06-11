package booking

import (
	"github.com/AriqF1/travel-web/internal/schedule"
	"github.com/AriqF1/travel-web/internal/user"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	UserID     		uint
	ScheduleID 		uint

	PassengerName  	string `json:"passenger_name" binding:"required"`
	SeatNumber     	string `json:"seat_number" binding:"required"`

	User     		user.User     `gorm:"foreignKey:UserID"`
	Schedule 		schedule.Schedule `gorm:"foreignKey:ScheduleID"`
}