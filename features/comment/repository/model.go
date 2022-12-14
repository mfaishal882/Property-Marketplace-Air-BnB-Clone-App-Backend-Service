package repository

import (
	_comment "api-airbnb-alta/features/comment"

	"gorm.io/gorm"
)

//struct gorm model
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

//mengubah struct core ke struct model gorm
func fromCore(dataCore _comment.Core) Comment {
	modelData := Comment{
		Title:      dataCore.Title,
		Comment:    dataCore.Comment,
		Rating:     dataCore.Rating,
		UserID:     dataCore.UserID,
		PropertyID: dataCore.PropertyID,
	}
	return modelData
}

//mengubah struct model gorm ke struct core
func (dataModel *Comment) toCore() _comment.Core {
	return _comment.Core{
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

//mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Comment) []_comment.Core {
	var dataCore []_comment.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
