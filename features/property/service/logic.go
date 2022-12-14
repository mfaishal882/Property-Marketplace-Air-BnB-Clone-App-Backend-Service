package service

import (
	"api-airbnb-alta/features/property"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"api-airbnb-alta/utils/thirdparty"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type propertyService struct {
	propertyRepository property.RepositoryInterface
	validate           *validator.Validate
}

func New(repo property.RepositoryInterface) property.ServiceInterface {
	return &propertyService{
		propertyRepository: repo,
		validate:           validator.New(),
	}
}

// Create implements properties.ServiceInterface
func (service *propertyService) Create(input property.Core, c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	input.UserID = uint(id)
	input.RatingAverage = 0
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.UploadProfile(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(res)
		input.ImageThumbnailUrl = res
	} else {
		input.ImageThumbnailUrl = "https://thumbs.dreamstime.com/b/house-vector-icon-home-logo-isolated-white-background-house-vector-icon-home-logo-vector-illustration-138343234.jpg"
	}

	errCreate := service.propertyRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements properties.ServiceInterface
func (service *propertyService) GetAll(queryName, queryCity, queryPropertyType string) (data []property.Core, err error) {

	if queryName == "" && queryCity == "" && queryPropertyType == "" {
		data, err = service.propertyRepository.GetAll()
	} else {
		data, err = service.propertyRepository.GetAllWithSearch(queryName, queryCity, queryPropertyType)
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	fmt.Println("Isi result get all property ,", data)

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, nil
}

// GetById implements properties.ServiceInterface
func (service *propertyService) GetById(id int) (data property.Core, err error) {
	data, err = service.propertyRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return property.Core{}, helper.ServiceErrorMsg(err)
	}

	if data == (property.Core{}) {
		helper.LogDebug("Get data success. No data.")
		return property.Core{}, errors.New("Get data success. No data.")
	}

	return data, err
}

// Delete implements property.ServiceInterface
func (service *propertyService) Delete(id int) error {
	err := service.propertyRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

// Update implements property.ServiceInterface
func (service *propertyService) Update(input property.Core, id int) error {
	err := service.propertyRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetPropertyImages implements property.ServiceInterface
func (service *propertyService) GetPropertyImages(id int) (data []property.PropertyImage, err error) {
	data, err = service.propertyRepository.GetPropertyImages(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

// GetPropertyComments implements property.ServiceInterface
func (service *propertyService) GetPropertyComments(id int) (data []property.Comment, err error) {
	data, err = service.propertyRepository.GetPropertyComments(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

// GetAvailbility implements property.ServiceInterface
func (service *propertyService) GetAvailbility(id uint, checkinDate time.Time, checkoutDate time.Time) (result string, err error) {

	if checkinDate.After(checkoutDate) || checkinDate.Equal(checkoutDate) {
		log.Error("Check In Date must before Checkout Date.")
		return result, errors.New("Check In Date must before Checkout Date.")
	}

	result, err = service.propertyRepository.GetAvailability(id, checkinDate, checkoutDate)
	if err != nil {
		log.Error(err.Error())
		return result, helper.ServiceErrorMsg(err)
	}
	return result, err
}
