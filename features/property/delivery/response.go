package delivery

import (
	property "api-airbnb-alta/features/property"
)

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

type PropertyResponseGetAll struct {
	ID                uint    `json:"id"`
	PropertyName      string  `json:"property_name" form:"property_name"`
	PricePerNight     int     `json:"price_per_night" form:"price_per_night"`
	City              string  `json:"city" form:"city"`
	Facilities        string  `json:"facilities" form:"facilities"`
	RatingAverage     float64 `json:"rating_average" form:"rating_average"`
	ImageThumbnailUrl string  `json:"image_thumbnail_url" form:"image_thumbnail_url"`
}

func fromCore(dataCore property.Core) PropertyResponseGetAll {
	return PropertyResponseGetAll{
		ID:                dataCore.ID,
		PropertyName:      dataCore.PropertyName,
		PricePerNight:     dataCore.PricePerNight,
		City:              dataCore.City,
		Facilities:        dataCore.Facilities,
		RatingAverage:     dataCore.RatingAverage,
		ImageThumbnailUrl: dataCore.ImageThumbnailUrl,
	}
}

func fromCoreList(dataCore []property.Core) []PropertyResponseGetAll {
	var dataResponse []PropertyResponseGetAll
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreById(dataCore property.Core) PropertyResponse {
	return PropertyResponse{
		ID:                dataCore.ID,
		PropertyName:      dataCore.PropertyName,
		PricePerNight:     dataCore.PricePerNight,
		Description:       dataCore.Description,
		Address:           dataCore.Address,
		City:              dataCore.City,
		ContactNumber:     dataCore.ContactNumber,
		Facilities:        dataCore.Facilities,
		PropertyType:      dataCore.PropertyType,
		RatingAverage:     dataCore.RatingAverage,
		ImageThumbnailUrl: dataCore.ImageThumbnailUrl,
		UserID:            dataCore.UserID,
	}
}
