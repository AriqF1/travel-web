package booking

type BookingRequest struct {
	ScheduleID 		uint   `json:"schedule_id" binding:"required"`
	PassengerName  	string `json:"passenger_name" binding:"required"`
	SeatNumber     	string `json:"seat_number" binding:"required"`	
}
