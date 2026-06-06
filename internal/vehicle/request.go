package vehicle


type CreateVehicleRequest struct {
	Name 			string `json:"name" binding:"required"`
	PlateNumber 	string `json:"plate_number" binding:"required"`
	SeatCount 		int    `json:"seat_count" binding:"required"`
}

type GetVehicleByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdateVehicleRequest struct {
	Name 			string `json:"name" binding:"required"`
	PlateNumber 	string `json:"plate_number" binding:"required"`
	SeatCount 		int    `json:"seat_count" binding:"required"`
}