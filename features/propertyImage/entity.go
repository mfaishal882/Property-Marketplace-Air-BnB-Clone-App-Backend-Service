package propertyImage

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Title      string `validate:"required"`
	ImageUrl   string
	PropertyID uint `validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Property struct {
	ID                uint
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
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, c echo.Context) error
	Delete(id int) error
	GetPropertyById(id int) (data Property, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetPropertyById(id int) (data Property, err error)
}
