package repository

import (
	"api-airbnb-alta/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint
	FullName        string
	Email           string `valiidate:"required,email"`
	Password        string `valiidate:"required"`
	Phone           string
	Gender          string
	ProfileImageUrl string
	IsHosting       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

//DTO

func (dataModel User) toCore() auth.Core {
	return auth.Core{
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
