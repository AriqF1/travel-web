package user

import (
	"github.com/AriqF1/travel-web/pkg/database"
)

func GetUserByID(id uint)(User, error) {
	var user User

	err := database.DB.First(&user, id).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}