package booking

import (
	"errors"
	"gorm.io/gorm"

	"github.com/AriqF1/travel-web/internal/schedule"
	"github.com/AriqF1/travel-web/internal/user"

	"github.com/AriqF1/travel-web/pkg/database"
)

func CreateBooking(userID uint, req BookingRequest) error {
	_, err := schedule.GetScheduleByID(req.ScheduleID)

	if err != nil {
		return errors.New("schedule not found")
	}

	_, err = user.GetUserByID(userID)

	if err != nil {
		return errors.New("user not found")
	}

	err = database.DB.
	Where(
		"schedule_id = ? AND seat_number = ?",
		req.ScheduleID,
		req.SeatNumber,
	).
	First(&Booking{}).
	Error

	if err == nil {
		return errors.New("seat already booked")
	}

	if err != gorm.ErrRecordNotFound {
		return err
	}

	newBooking := Booking{
		UserID: userID,
		ScheduleID: req.ScheduleID,
		PassengerName: req.PassengerName,
		SeatNumber: req.SeatNumber,
	}

	return database.DB.Create(&newBooking).Error
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