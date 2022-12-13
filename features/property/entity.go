package property

import (
	"api-airbnb-alta/features/user"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	PropertyName   string `valiidate:"required"`
	PricePerNight  int    `valiidate:"required"`
	Description    string
	Address        string `valiidate:"required"`
	City           string `valiidate:"required"`
	ContactNumber  string
	Fasilities     string
	PropertyType   string
	RatingAverage  float64
	ImageThumbnail string
	UserID         uint
	User           user.Core
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ServiceInterface interface {
	GetAll(queryName, queryCity, queryPropertyType string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	// GetAll() (data []Core, err error)
	GetAllWithSearch(queryName, queryCity, queryPropertyType string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}
