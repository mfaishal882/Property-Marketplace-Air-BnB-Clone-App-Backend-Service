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
	Facilities     string `json:"facilities" form:"facilities"`
	PropertyType   string `json:"property_type" form:"property_type"`
	ImageThumbnail string `json:"image_thumbnail" form:"image_thumbnail"`
	UserID         uint   `json:"user_id" form:"user_id"`
}

type CheckAvailRequest struct {
	CheckinDate  string `json:"checkin_date" form:"checkin_date"`
	CheckoutDate int    `json:"checkout_date" form:"checkout_date"`
	PropertyID   string `json:"property_id" form:"property_id"`
}

func toCore(propertyInput InsertRequest) property.Core {
	propertyCoreData := property.Core{
		PropertyName:      propertyInput.PropertyName,
		PricePerNight:     propertyInput.PricePerNight,
		Description:       propertyInput.Description,
		Address:           propertyInput.Address,
		City:              propertyInput.City,
		ContactNumber:     propertyInput.ContactNumber,
		Facilities:        propertyInput.Facilities,
		PropertyType:      propertyInput.PropertyType,
		ImageThumbnailUrl: propertyInput.ImageThumbnail,
		UserID:            propertyInput.UserID,
	}
	return propertyCoreData
}
