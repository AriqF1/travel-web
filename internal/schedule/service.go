package schedule

import (
	"time"
	"errors"

	"github.com/AriqF1/travel-web/internal/vehicle"
	"github.com/AriqF1/travel-web/pkg/database"
)

func CreateSchedule(req ScheduleRequest) error {
	_, err := vehicle.GetVehicleByID(req.VehicleID)

	if err != nil {
		return errors.New("vehicle not found")
	}

	departureTime, err := time.Parse(
		time.RFC3339,
		req.DepartureTime,
	)

	if err != nil {
		return err
	}

	newSchedule := Schedule{
		VehicleID: req.VehicleID,
		Origin: req.Origin,
		Destination: req.Destination,
		DepartureTime: departureTime,
		Price: req.Price,
	}

	return database.DB.Create(&newSchedule).Error
}

func GetSchedules()([]Schedule, error) {
	var schedules []Schedule

	err := database.DB.Preload("Vehicle").Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, err
}

func GetScheduleByID(id uint)(Schedule, error) {
	var schedule Schedule

	err := database.DB.Preload("Vehicle").First(&schedule, id).Error

	if err != nil {
		return Schedule{}, err
	}

	return schedule, nil
}

func UpdateSchedule(id uint, req ScheduleRequest) error {
	schedule, err := GetScheduleByID(id)

	if err != nil {
		return err
	}

	departureTime, err := time.Parse(
		time.RFC3339,
		req.DepartureTime,
	)

	if err != nil {
		return err
	}

	schedule.VehicleID = req.VehicleID
	schedule.Origin = req.Origin
	schedule.Destination = req.Destination
	schedule.DepartureTime = departureTime
	schedule.Price = req.Price

	return database.DB.Save(&schedule).Error
}

func DeleteSchedule(id uint) error {
	schedule, err := GetScheduleByID(id)

	if err != nil {
		return err
	}
	
	return database.DB.Delete(&schedule).Error
}