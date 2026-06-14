package booking

import (
	"time"

	"github.com/AriqF1/travel-web/internal/user"
	"github.com/AriqF1/travel-web/internal/vehicle"
)

type BookingResponse struct {
	ID            uint             			`json:"id"`
	PassengerName string            		`json:"passenger_name"`
	SeatNumber    string              		`json:"seat_number"`
	
	Origin        string            		`json:"origin"`
	Destination   string            		`json:"destination"`
	DepartureTime string            		`json:"departure_time"`

	Vehicle        vehicle.VehicleResponse 	`json:"vehicle"`
	BookedBy       user.UserResponse  		`json:"booked_by"`
}

func ToBookingResponse(
	b Booking,
) BookingResponse {

	return BookingResponse{
		ID:            		b.ID,
		PassengerName: 		b.PassengerName,
		SeatNumber: 		b.SeatNumber,
		Origin: 			b.Schedule.Origin,
		Destination: 		b.Schedule.Destination,
		DepartureTime: 		b.Schedule.DepartureTime.Format(time.RFC3339),
		Vehicle: 			vehicle.ToVehicleResponse(b.Schedule.Vehicle),
		BookedBy: 			user.ToUserResponse(b.User),
	}
}

func ToBookingResponses(
	bookings []Booking,
) []BookingResponse {
	 
	var responses []BookingResponse

	for _, booking := range bookings {

		responses = append(
			responses,
			ToBookingResponse(booking),
		)
	}

	return responses
}