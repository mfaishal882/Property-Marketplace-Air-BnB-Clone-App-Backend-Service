package comment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Title      string  `valiidate:"required"`
	Comment    string  `valiidate:"required"`
	Rating     float64 `valiidate:"required"`
	UserID     uint
	PropertyID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, c echo.Context) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}
