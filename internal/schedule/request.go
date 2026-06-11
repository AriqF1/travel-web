package schedule

type ScheduleRequest struct {
	VehicleID 		uint   `json:"vehicle_id" binding:"required"`
	Origin    		string `json:"origin" binding:"required"`
	Destination 	string `json:"destination" binding:"required"`
	DepartureTime 	string `json:"departure_time" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Price       	int    `json:"price" binding:"required,min=0"`
}
