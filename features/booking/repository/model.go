package repository

import (
	"api-airbnb-alta/features/booking"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	CheckinDate   time.Time
	CheckoutDate  time.Time
	PricePerNight float64
	GrossAmount   float64
	BookingStatus string
	User          User
	Property      Property
	UserID        uint
	PropertyID    uint
}

type Property struct {
	gorm.Model
	PropertyName      string  `valiidate:"required"`
	PricePerNight     float64 `valiidate:"required"`
	Description       string
	Address           string `valiidate:"required"`
	City              string `valiidate:"required"`
	ContactNumber     string
	Facilities        string
	PropertyType      string
	RatingAverage     float64
	ImageThumbnailUrl string
	UserID            uint
	User              User
	Bookings          []Booking
}

type User struct {
	gorm.Model
	FullName        string
	Email           string
	Password        string
	Phone           string
	Gender          string
	ProfileImageUrl string
	IsHosting       string
	Bookings        []Booking
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore booking.Core) Booking {
	modelData := Booking{
		CheckinDate:   dataCore.CheckinDate,
		CheckoutDate:  dataCore.CheckoutDate,
		PricePerNight: dataCore.PricePerNight,
		GrossAmount:   dataCore.GrossAmount,
		BookingStatus: dataCore.BookingStatus,
		UserID:        dataCore.UserID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Booking) toCore() booking.Core {
	return booking.Core{
		ID:            dataModel.ID,
		CheckinDate:   dataModel.CheckinDate,
		CheckoutDate:  dataModel.CheckoutDate,
		PricePerNight: dataModel.PricePerNight,
		GrossAmount:   dataModel.GrossAmount,
		BookingStatus: dataModel.BookingStatus,
		UserID:        dataModel.UserID,
		User: booking.User{
			ID:       dataModel.User.ID,
			FullName: dataModel.User.FullName,
		},
		Property: booking.Property{
			ID:            dataModel.Property.ID,
			PropertyName:  dataModel.Property.PropertyName,
			PricePerNight: dataModel.Property.PricePerNight,
		},
		PropertyID: dataModel.PropertyID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Booking) []booking.Core {
	var dataCore []booking.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
