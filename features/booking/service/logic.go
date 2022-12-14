package service

import (
	"api-airbnb-alta/features/booking"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type bookingService struct {
	bookingRepository booking.RepositoryInterface
	validate          *validator.Validate
}

func New(repo booking.RepositoryInterface) booking.ServiceInterface {
	return &bookingService{
		bookingRepository: repo,
		validate:          validator.New(),
	}
}

// Create implements booking.ServiceInterface
func (service *bookingService) Create(input booking.Core, c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	// userName := middlewares.ExtractTokenUserName(c)
	input.UserID = uint(id)
	// input.User.FullName = userName

	errCreate := service.bookingRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements booking.ServiceInterface
func (service *bookingService) GetAll() (data []booking.Core, err error) {
	data, err = service.bookingRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}
	return data, nil
}

// GetById implements booking.ServiceInterface
func (service *bookingService) GetById(id int) (data booking.Core, err error) {
	data, err = service.bookingRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if data == (booking.Core{}) {
		helper.LogDebug("Get data success. No data.")
		return booking.Core{}, errors.New("Get data success. No data.")
	}

	return data, err

}
