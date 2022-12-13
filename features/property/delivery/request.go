package delivery

import (
	property "api-airbnb-alta/features/property"
)

type InsertRequest struct {
	PropertyName   string `json:"property_name" form:"property_name"`
	PricePerNight  int    `json:"price_per_night" form:"price_per_night"`
	Description    string `json:"description" form:"description"`
	Address        string `json:"address" form:"address"`
	City           string `json:"city" form:"city"`
	ContactNumber  string `json:"contact_number" form:"contact_number"`
	Fasilities     string `json:"fasilities" form:"fasilities"`
	PropertyType   string `json:"property_type" form:"property_type"`
	ImageThumbnail string `json:"image_thumbnail" form:"image_thumbnail"`
	UserID         uint   `json:"user_id" form:"user_id"`
}

func toCore(propertyInput InsertRequest) property.Core {
	propertyCoreData := property.Core{
		PropertyName:   propertyInput.PropertyName,
		Description:    propertyInput.Description,
		Address:        propertyInput.Address,
		City:           propertyInput.City,
		ContactNumber:  propertyInput.ContactNumber,
		Fasilities:     propertyInput.Fasilities,
		PropertyType:   propertyInput.PropertyType,
		ImageThumbnail: propertyInput.ImageThumbnail,
		UserID:         propertyInput.UserID,
	}
	return propertyCoreData
}
