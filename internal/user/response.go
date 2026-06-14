package user


type UserResponse struct {
	ID 	 	uint   		`json:"id"`
	Name 	string 		`json:"name"`
	Email 	string 		`json:"email"`
}

func ToUserResponse(
	u User,
) UserResponse {

	return UserResponse{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}

func ToUserResponses(
	users []User,
) []UserResponse {

	var responses []UserResponse

	for _, user := range users {

		responses = append(
			responses,
			ToUserResponse(user),
		)
	}
	return responses
}