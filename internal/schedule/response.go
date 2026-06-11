package schedule

import (
	"time"

	"github.com/AriqF1/travel-web/internal/vehicle"
)

type ScheduleResponse struct {
	ID            uint            			`json:"id"`
	Origin        string          			`json:"origin"`
	Destination   string          			`json:"destination"`
	DepartureTime string          			`json:"departure_time"`
	Price         int             			`json:"price"`
	Vehicle       vehicle.VehicleResponse 	`json:"vehicle"`
}

func ToScheduleResponse(
	s Schedule,
) ScheduleResponse {

	return ScheduleResponse{
		ID:            s.ID,
		Origin:        s.Origin,
		Destination:   s.Destination,
		DepartureTime: s.DepartureTime.Format(time.RFC3339),
		Price:         s.Price,

		Vehicle: vehicle.ToVehicleResponse(s.Vehicle),
	}
}

func ToScheduleResponses(
	schedules []Schedule,
) []ScheduleResponse {

	var responses []ScheduleResponse

	for _, schedule := range schedules {

		responses = append(
			responses,
			ToScheduleResponse(schedule),
		)
	}

	return responses
}