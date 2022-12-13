package delivery

import (
	property "api-airbnb-alta/features/property"
)

type PropertyResponse struct {
	ID             uint    `json:"id"`
	PropertyName   string  `json:"property_name" form:"property_name"`
	PricePerNight  int     `json:"price_per_night" form:"price_per_night"`
	Description    string  `json:"description" form:"description"`
	Address        string  `json:"address" form:"address"`
	City           string  `json:"city" form:"city"`
	ContactNumber  string  `json:"contact_number" form:"contact_number"`
	Fasilities     string  `json:"fasilities" form:"fasilities"`
	PropertyType   string  `json:"property_type" form:"property_type"`
	RatingAverage  float64 `json:"rating_average" form:"rating_average"`
	ImageThumbnail string  `json:"image_thumbnail" form:"image_thumbnail"`
	UserID         uint    `json:"user_id" form:"user_id"`
}

func fromCore(dataCore property.Core) PropertyResponse {
	return PropertyResponse{
		ID:             dataCore.ID,
		PropertyName:   dataCore.PropertyName,
		Description:    dataCore.Description,
		Address:        dataCore.Address,
		City:           dataCore.City,
		ContactNumber:  dataCore.ContactNumber,
		Fasilities:     dataCore.Fasilities,
		PropertyType:   dataCore.PropertyType,
		RatingAverage:  dataCore.RatingAverage,
		ImageThumbnail: dataCore.ImageThumbnail,
		UserID:         dataCore.UserID,
	}
}

func fromCoreList(dataCore []property.Core) []PropertyResponse {
	var dataResponse []PropertyResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
