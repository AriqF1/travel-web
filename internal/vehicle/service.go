package vehicle

import "github.com/AriqF1/travel-web/pkg/database"

func createVehicle(req CreateVehicleRequest) error {

	vehicle := Vehicle{
		Name:        req.Name,
		PlateNumber: req.PlateNumber,
		SeatCount:   req.SeatCount,
	}

	return database.DB.Create(&vehicle).Error
}

func GetVehicle()([]Vehicle, error) {
	var vehicles []Vehicle

	err := database.DB.Find(&vehicles).Error
	
	if err != nil {
		return nil, err
	}

	return vehicles, err	
}

func GetVehicleByID(id uint)(Vehicle, error) {
	var vehicle Vehicle

	err := database.DB.First(&vehicle, id).Error

	if err != nil {
		return Vehicle{}, err
	}

	return vehicle, nil
}

func DeleteVehicle(id uint) error {
	vehicle, err := GetVehicleByID(id)

	if err != nil {
		return err
	}

	return database.DB.Delete(&vehicle).Error

}

func UpdateVehicle(id uint, req UpdateVehicleRequest) error {
	vehicle, err := GetVehicleByID(id)

	if err != nil {
		return err
	}

	vehicle.Name = req.Name
	vehicle.PlateNumber = req.PlateNumber
	vehicle.SeatCount = req.SeatCount

	return database.DB.Save(&vehicle).Error
}