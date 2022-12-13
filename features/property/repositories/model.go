package repositories

import (
	_property "api-airbnb-alta/features/property"

	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	PropertyName   string `valiidate:"required"`
	PricePerNight  int    `valiidate:"required"`
	Description    string
	Address        string `valiidate:"required"`
	City           string `valiidate:"required"`
	ContactNumber  string
	GuestCapacity  string
	Fasilities     string
	BookingStatus  string
	PropertyType   string
	RatingAverage  float64
	ImageThumbnail string
	UserID         uint
	User           User
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
}

func fromCore(dataCore _property.Core) Property {
	userGorm := Property{
		PropertyName:   dataCore.PropertyName,
		Description:    dataCore.Description,
		Address:        dataCore.Address,
		City:           dataCore.City,
		ContactNumber:  dataCore.ContactNumber,
		Fasilities:     dataCore.Fasilities,
		PropertyType:   dataCore.PropertyType,
		ImageThumbnail: dataCore.ImageThumbnail,
		UserID:         dataCore.UserID,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Property) toCore() _property.Core {
	return _property.Core{
		ID:             dataModel.ID,
		PropertyName:   dataModel.PropertyName,
		Description:    dataModel.Description,
		Address:        dataModel.Address,
		City:           dataModel.City,
		ContactNumber:  dataModel.ContactNumber,
		Fasilities:     dataModel.Fasilities,
		PropertyType:   dataModel.PropertyType,
		RatingAverage:  dataModel.RatingAverage,
		ImageThumbnail: dataModel.ImageThumbnail,
		UserID:         dataModel.UserID,
		CreatedAt:      dataModel.CreatedAt,
		UpdatedAt:      dataModel.UpdatedAt,
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
