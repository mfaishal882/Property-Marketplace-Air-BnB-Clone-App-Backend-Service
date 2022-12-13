package repository

import (
	_propertyImage "api-airbnb-alta/features/propertyImage"

	"gorm.io/gorm"
)

//struct gorm model
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
}

// DTO
// mapping

//mengubah struct core ke struct model gorm
func fromCore(dataCore _propertyImage.Core) PropertyImage {
	modelData := PropertyImage{
		Title:      dataCore.Title,
		ImageUrl:   dataCore.ImageUrl,
		PropertyID: dataCore.PropertyID,
	}
	return modelData
}

//mengubah struct model gorm ke struct core
func (dataModel *PropertyImage) toCore() _propertyImage.Core {
	return _propertyImage.Core{
		ID:         dataModel.ID,
		Title:      dataModel.Title,
		ImageUrl:   dataModel.ImageUrl,
		PropertyID: dataModel.PropertyID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

//mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []PropertyImage) []_propertyImage.Core {
	var dataCore []_propertyImage.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
