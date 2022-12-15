package service

import (
	"api-airbnb-alta/features/booking"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"errors"
	"fmt"

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
	input.UserID = uint(id)

	availStatus, errAvail := service.bookingRepository.GetAvailability(input.PropertyID, input.CheckinDate, input.CheckoutDate)
	if errAvail != nil {
		return errors.New("Failed check checkin checkout availbility.")
	}

	fmt.Println("\n\nprop id", input.Property.ID)
	fmt.Println("\n\n checkin", input.CheckinDate)
	fmt.Println("\n\n checkout", input.CheckoutDate)
	fmt.Println("\n\navail status", availStatus)

	if availStatus != "Available" {
		return errors.New("Check In and Check Out date not available. Please pick another one.")
	}

	errCreate := service.bookingRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements booking.ServiceInterface
func (service *bookingService) GetAll(userId int) (data []booking.Core, err error) {
	data, err = service.bookingRepository.GetAll(userId)

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
func (service *bookingService) GetById(id int, userId int) (data booking.Core, err error) {
	data, err = service.bookingRepository.GetById(id, userId)
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
