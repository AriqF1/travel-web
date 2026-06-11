package vehicle

type VehicleResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PlateNumber string `json:"plate_number"`
	SeatCount   int    `json:"seat_count"`
}

func ToVehicleResponse(
	v Vehicle,
) VehicleResponse {

	return VehicleResponse{
		ID:          v.ID,
		Name:        v.Name,
		PlateNumber: v.PlateNumber,
		SeatCount:   v.SeatCount,
	}
}

func ToVehicleResponses(
	vehicles []Vehicle,
) []VehicleResponse {

	var responses []VehicleResponse

	for _, vehicle := range vehicles {

		responses = append(
			responses,
			ToVehicleResponse(vehicle),
		)
	}

	return responses
}