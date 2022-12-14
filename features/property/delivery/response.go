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

type PropertyImagesResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" form:"title"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	PropertyID uint   `json:"properties_id" form:"properties_id"`
}

type PropertyCommentResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title" form:"title"`
	Comment    string  `json:"comment" form:"comment"`
	Rating     float64 `json:"rating" form:"rating"`
	UserID     uint    `json:"user_id" form:"user_id"`
	PropertyID uint    `json:"properties_id" form:"properties_id"`
}

type PropertyResponseCheckAvailbility struct {
	PropertyID    uint   `json:"property_id"`
	PropertyName  string `json:"property_name"`
	CheckinDate   string `json:"checkin_date"`
	CheckoutDate  string `json:"checkout_date"`
	BookingStatus string `json:"booking_status"`
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

func fromPropertyImages(dataCore property.PropertyImage) PropertyImagesResponse {
	return PropertyImagesResponse{
		ID:         dataCore.ID,
		Title:      dataCore.Title,
		ImageUrl:   dataCore.ImageUrl,
		PropertyID: dataCore.PropertyID,
	}
}

func fromPropertyImagesList(dataCore []property.PropertyImage) []PropertyImagesResponse {
	var dataResponse []PropertyImagesResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromPropertyImages(v))
	}
	return dataResponse
}

func fromPropertyComment(dataCore property.Comment) PropertyCommentResponse {
	return PropertyCommentResponse{
		ID:         dataCore.ID,
		Title:      dataCore.Title,
		Comment:    dataCore.Comment,
		Rating:     dataCore.Rating,
		UserID:     dataCore.UserID,
		PropertyID: dataCore.PropertyID,
	}
}

func fromPropertyCommentList(dataCore []property.Comment) []PropertyCommentResponse {
	var dataResponse []PropertyCommentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromPropertyComment(v))
	}
	return dataResponse
}

func fromCoreCheckAvailbility(dataCore property.Core, checkinDate string, checkoutDate string, bookingStatus string) PropertyResponseCheckAvailbility {
	return PropertyResponseCheckAvailbility{
		PropertyID:    dataCore.ID,
		PropertyName:  dataCore.PropertyName,
		CheckinDate:   checkinDate,
		CheckoutDate:  checkoutDate,
		BookingStatus: bookingStatus,
	}
}
