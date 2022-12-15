package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	FullName        string `validate:"required"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required"`
	Phone           string `validate:"required"`
	Gender          string
	ProfileImageUrl string
	IsHosting       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Property struct {
	ID                uint
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
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetProperties(id int) (data []Property, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
	GetProperties(id int) (data []Property, err error)
}
