package booking

import (
	"api-airbnb-alta/features/property"
	"api-airbnb-alta/features/user"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	CheckinDate   time.Time
	CheckoutDate  time.Time
	PricePerNight float64
	GrossAmount   float64
	BookingStatus string
	UserID        uint
	PropertyID    uint
	User          user.Core
	Property      property.Core
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
}
