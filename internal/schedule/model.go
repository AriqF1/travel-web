package schedule

import (
	"time"

	"github.com/AriqF1/travel-web/internal/vehicle"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	VehicleID uint

	Origin 			string
	Destination 	string
	DepartureTime 	time.Time
	Price 			int

	Vehicle vehicle.Vehicle `gorm:"foreignKey:VehicleID"`
}