package repository

import (
	_user "api-airbnb-alta/features/user"

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
}

type Comment struct {
	gorm.Model
	Title      string
	Comment    string
	Rating     float64
	UserID     uint
	PropertyID uint
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
