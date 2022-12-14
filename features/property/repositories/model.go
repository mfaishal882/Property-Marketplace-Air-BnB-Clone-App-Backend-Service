package repositories

import (
	_property "api-airbnb-alta/features/property"

	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	PropertyName      string `valiidate:"required"`
	PricePerNight     int    `valiidate:"required"`
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
	PropertyImages    []PropertyImage
	Comments          []Comment
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
	Properties      []Property
	Comments        []Comment
}

type PropertyImage struct {
	gorm.Model
	Title      string
	ImageUrl   string `valiidate:"required"`
	PropertyID uint   `valiidate:"required"`
}

type Comment struct {
	gorm.Model
	Title      string
	Comment    string
	Rating     float64
	UserID     uint
	PropertyID uint
}

func fromCore(dataCore _property.Core) Property {
	userGorm := Property{
		PropertyName:      dataCore.PropertyName,
		PricePerNight:     dataCore.PricePerNight,
		Description:       dataCore.Description,
		Address:           dataCore.Address,
		City:              dataCore.City,
		ContactNumber:     dataCore.ContactNumber,
		Facilities:        dataCore.Facilities,
		PropertyType:      dataCore.PropertyType,
		ImageThumbnailUrl: dataCore.ImageThumbnailUrl,
		UserID:            dataCore.UserID,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Property) toCore() _property.Core {
	return _property.Core{
		ID:                dataModel.ID,
		PropertyName:      dataModel.PropertyName,
		PricePerNight:     dataModel.PricePerNight,
		Description:       dataModel.Description,
		Address:           dataModel.Address,
		City:              dataModel.City,
		ContactNumber:     dataModel.ContactNumber,
		Facilities:        dataModel.Facilities,
		PropertyType:      dataModel.PropertyType,
		RatingAverage:     dataModel.RatingAverage,
		ImageThumbnailUrl: dataModel.ImageThumbnailUrl,
		UserID:            dataModel.UserID,
		CreatedAt:         dataModel.CreatedAt,
		UpdatedAt:         dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Property) []_property.Core {
	var dataCore []_property.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *PropertyImage) toCorePropertyImage() _property.PropertyImage {
	return _property.PropertyImage{
		ID:         dataModel.ID,
		Title:      dataModel.Title,
		ImageUrl:   dataModel.ImageUrl,
		PropertyID: dataModel.PropertyID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

func toPropertyImagesList(dataModel []PropertyImage) []_property.PropertyImage {
	var dataCore []_property.PropertyImage
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCorePropertyImage())
	}
	return dataCore
}

func (dataModel *Comment) toCorePropertyComment() _property.Comment {
	return _property.Comment{
		ID:         dataModel.ID,
		Title:      dataModel.Title,
		Comment:    dataModel.Comment,
		Rating:     dataModel.Rating,
		UserID:     dataModel.UserID,
		PropertyID: dataModel.PropertyID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

func toPropertyCommentList(dataModel []Comment) []_property.Comment {
	var dataCore []_property.Comment
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCorePropertyComment())
	}
	return dataCore
}
