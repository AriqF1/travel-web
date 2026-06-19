package schedule

import (
	"time"

	"github.com/AriqF1/travel-web/internal/vehicle"
)

type ScheduleResponse struct {
	ID            	uint            			`json:"id"`
	Origin        	string          			`json:"origin"`
	Destination   	string          			`json:"destination"`
	DepartureTime 	string          			`json:"departure_time"`
	Price         	int             			`json:"price"`
	AvailableSeats 	int 						`json:"available_seats"`

	Vehicle       	vehicle.VehicleResponse 	`json:"vehicle"`
}

func ToScheduleResponse(
	s Schedule,
	availableSeats int,
) ScheduleResponse {

	return ScheduleResponse{
		ID				: s.ID,
		Origin			: s.Origin,
		Destination		: s.Destination,
		DepartureTime	: s.DepartureTime.Format(time.RFC3339),
		Price			: s.Price,
		AvailableSeats 	: availableSeats,

		Vehicle: vehicle.ToVehicleResponse(s.Vehicle),
	}
}

func ToScheduleResponses(
	schedules []Schedule,
) []ScheduleResponse {

	var responses []ScheduleResponse

	for _, schedule := range schedules {
		
		availableSeats, _ := GetAvailableSeats(schedule.ID)

		responses = append(
			responses,
			ToScheduleResponse(schedule, availableSeats),
		)		
	}

	return responses
}