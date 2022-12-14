package repository

import (
	_user "api-airbnb-alta/features/user"
	"time"

	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	FullName        string `validate:"required"`
	Email           string `validate:"required,email,unique"`
	Password        string `validate:"required"`
	Phone           string `validate:"required"`
	Gender          string
	ProfileImageUrl string
	IsHosting       string
	Propertys       []Property
	Comments        []Comment
	Booking         []Booking
}

type PropertyImage struct {
	gorm.Model
	Title      string
	ImageUrl   string `valiidate:"required"`
	PropertyID uint   `valiidate:"required"`
}

type Property struct {
	gorm.Model
	PropertyName      string
	PricePerNight     int
	Description       string
	Address           string
	City              string
	ContactNumber     string
	Facilities        string
	PropertyType      string
	RatingAverage     float64
	ImageThumbnailUrl string
	UserID            uint
	PropertyImages    []PropertyImage
	Comments          []Comment
	Bookings          []Booking
}

type Comment struct {
	gorm.Model
	Title      string
	Comment    string
	Rating     float64
	UserID     uint
	PropertyID uint
}

type Booking struct {
	gorm.Model
	CheckinDate   time.Time
	CheckoutDate  time.Time
	PricePerNight float64
	GrossAmount   float64
	BookingStatus string
	UserID        uint
	PropertyID    uint
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _user.Core) User {
	userGorm := User{
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
		ProfileImageUrl: dataCore.ProfileImageUrl,
		IsHosting:       dataCore.IsHosting,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:              dataModel.ID,
		FullName:        dataModel.FullName,
		Email:           dataModel.Email,
		Password:        dataModel.Password,
		Phone:           dataModel.Phone,
		Gender:          dataModel.Gender,
		ProfileImageUrl: dataModel.ProfileImageUrl,
		IsHosting:       dataModel.IsHosting,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *Property) toCoreProperties() _user.Property {
	return _user.Property{
		ID:                dataModel.ID,
		PropertyName:      dataModel.PropertyName,
		PricePerNight:     dataModel.PricePerNight,
		Description:       dataModel.Address,
		Address:           dataModel.Address,
		City:              dataModel.City,
		ContactNumber:     dataModel.ContactNumber,
		Facilities:        dataModel.Facilities,
		PropertyType:      dataModel.PropertyType,
		RatingAverage:     dataModel.RatingAverage,
		ImageThumbnailUrl: dataModel.ImageThumbnailUrl,
		CreatedAt:         dataModel.CreatedAt,
		UpdatedAt:         dataModel.UpdatedAt,
	}
}

func toPropertiesList(dataModel []Property) []_user.Property {
	var dataCore []_user.Property
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreProperties())
	}
	return dataCore
}
