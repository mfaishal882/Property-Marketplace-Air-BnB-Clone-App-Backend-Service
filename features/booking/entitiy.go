package booking

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	CheckinDate   time.Time `validate:"required"`
	CheckoutDate  time.Time `validate:"required"`
	PricePerNight float64
	GrossAmount   float64
	BookingStatus string
	UserID        uint
	PropertyID    uint `validate:"required"`
	User          User
	Property      Property
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type User struct {
	ID       uint
	FullName string
}

type Property struct {
	ID            uint
	PropertyName  string
	PricePerNight float64
}
type ServiceInterface interface {
	GetAll(userId int) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int, userId int) (data Core, err error)
}

type RepositoryInterface interface {
	GetAll(userId int) (data []Core, err error)
	Create(input Core) error
	GetById(id int, userId int) (data Core, err error)
	GetAvailability(propertyId uint, checkinDate time.Time, checkoutData time.Time) (result string, err error)
}
