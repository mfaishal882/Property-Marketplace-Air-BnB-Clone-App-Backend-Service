package repository

import (
	"api-airbnb-alta/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	FullName  string
	Email     string `validate:"required,email"`
	Password  string `valudate:"required"`
	Team      string
	Role      string `valudate:"required"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//DTO

func (data User) toCore() auth.Core {
	return auth.Core{
		ID:        data.ID,
		FullName:  data.FullName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
