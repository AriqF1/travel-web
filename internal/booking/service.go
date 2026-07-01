package booking

import (
	"errors"
	"gorm.io/gorm"

	"github.com/AriqF1/travel-web/internal/schedule"
	"github.com/AriqF1/travel-web/internal/user"

	"github.com/AriqF1/travel-web/pkg/database"
)

func CreateBooking(userID uint, req BookingRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {

		// Mengambil schedule beserta vehicle
		var sch schedule.Schedule

		err := tx.
			Preload("Vehicle").
			First(&sch, req.ScheduleID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("schedule not found")
			}

			return err
		}

		// Validasi nomor kursi
		if err := ValidateSeatNumber(
			req.SeatNumber,
			sch.Vehicle.SeatCount,
		); err != nil {
			return err
		}

		// Mengecek apakah kursi sudah dibooking
		var existingBooking Booking

		err = tx.
			Where(
				"schedule_id = ? AND seat_number = ?",
				req.ScheduleID,
				req.SeatNumber,
			).
			First(&existingBooking).
			Error

		if err == nil {
			return errors.New("seat already booked")
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		newBooking := Booking{
			UserID:        userID,
			ScheduleID:    req.ScheduleID,
			PassengerName: req.PassengerName,
			SeatNumber:    req.SeatNumber,
		}

		if err := tx.Create(&newBooking).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetBookings()([]Booking, error) {
	var bookings []Booking

	err := database.DB.Preload("Schedule.Vehicle").Preload("User").Find(&bookings).Error

	if err != nil {
		return nil, err
	}

	return bookings, err
}

func GetBookingByID(id uint)(Booking, error){
	var booking Booking

	err := database.DB.Preload("Schedule.Vehicle").Preload("User").First(&booking, id).Error

	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}