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

type PropertyResponse struct {
	ID                uint    `json:"id"`
	PropertyName      string  `json:"property_name" form:"property_name"`
	PricePerNight     int     `json:"price_per_night" form:"price_per_night"`
	Description       string  `json:"description" form:"description"`
	Address           string  `json:"address" form:"address"`
	City              string  `json:"city" form:"city"`
	ContactNumber     string  `json:"contact_number" form:"contact_number"`
	Facilities        string  `json:"facilities" form:"facilities"`
	PropertyType      string  `json:"property_type" form:"property_type"`
	RatingAverage     float64 `json:"rating_average" form:"rating_average"`
	ImageThumbnailUrl string  `json:"image_thumbnail_url" form:"image_thumbnail_url"`
	UserID            uint    `json:"user_id" form:"user_id"`
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

func fromProperties(dataCore user.Property) PropertyResponse {
	return PropertyResponse{
		ID:                dataCore.ID,
		PropertyName:      dataCore.PropertyName,
		PricePerNight:     dataCore.PricePerNight,
		City:              dataCore.City,
		Facilities:        dataCore.Facilities,
		RatingAverage:     dataCore.RatingAverage,
		ImageThumbnailUrl: dataCore.ImageThumbnailUrl,
	}
}

func fromPropertiesList(dataCore []user.Property) []PropertyResponse {
	var dataResponse []PropertyResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromProperties(v))
	}
	return dataResponse
}
