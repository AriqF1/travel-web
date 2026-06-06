package vehicle

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Name 			string
	PlateNumber 	string
	SeatCount 		int
}