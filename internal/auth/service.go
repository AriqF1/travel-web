package auth

import (
	"errors"
	
	"github.com/AriqF1/travel-web/internal/user"
	"github.com/AriqF1/travel-web/pkg/database"
	"github.com/AriqF1/travel-web/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(req RegisterRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	newUser := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	return database.DB.Create(&newUser).Error
}

func Login(req LoginRequest) (string, error) {
	var user user.User

	// search user
	err := database.DB.
		Where("email = ?", req.Email).
		First(&user).Error

	if err != nil {
    	return "", err
	}
    // password check
	bcrypt.CompareHashAndPassword(
    []byte(user.Password),
    []byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}
    // generate token
	token, err := utils.GenerateToken(
    	user.ID,
    	user.Email,
	)

	if err != nil {
        return "", err
    }
    // return token
	return token, nil
}