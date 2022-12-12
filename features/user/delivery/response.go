package delivery

import "api-airbnb-alta/features/user"

type UserResponse struct {
	ID              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	ProfileImageUrl string `json:"profile_image_url"`
	IsHosting       string `json:"is_hosting"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:              dataCore.ID,
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
		ProfileImageUrl: dataCore.ProfileImageUrl,
		IsHosting:       dataCore.IsHosting,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
