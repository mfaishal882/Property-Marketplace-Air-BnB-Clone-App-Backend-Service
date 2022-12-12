package delivery

import "api-airbnb-alta/features/auth"

type UserResponse struct {
	ID              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	ProfileImageUrl string `json:"profile_image_url"`
	IsHosting       string `json:"is_hosting"`
	Token           string `json:"token"`
}

func FromCore(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:              dataCore.ID,
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
		ProfileImageUrl: dataCore.ProfileImageUrl,
		IsHosting:       dataCore.IsHosting,
		Token:           token,
	}
}
