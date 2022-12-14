package property

import (
	"api-airbnb-alta/features/user"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
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
	User              user.Core
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type PropertyImage struct {
	ID         uint
	Title      string
	ImageUrl   string `valiidate:"required"`
	PropertyID uint   `valiidate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Comment struct {
	ID         uint
	Title      string
	Comment    string
	Rating     float64
	UserID     uint
	PropertyID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceInterface interface {
	GetAll(queryName, queryCity, queryPropertyType string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetPropertyImages(id int) (data []PropertyImage, err error)
	GetPropertyComments(id int) (data []Comment, err error)
	GetAvailbility(id uint, checkinDate time.Time, checkoutDate time.Time) (result string, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(queryPropertyName, queryCity, queryPropertyType string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	GetPropertyImages(id int) (data []PropertyImage, err error)
	GetPropertyComments(id int) (data []Comment, err error)
	GetAvailability(propertyId uint, checkinDate time.Time, checkoutData time.Time) (result string, err error)
}
