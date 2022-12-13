package service

import (
	"api-airbnb-alta/features/property"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"errors"

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
	errCreate := service.propertyRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements properties.ServiceInterface
func (service *propertyService) GetAll(queryName, queryCity, queryPropertyType string) (data []property.Core, err error) {

	data, err = service.propertyRepository.GetAllWithSearch(queryName, queryCity, queryPropertyType)
	if err != nil {
		return nil, err
	}
	return
}

// GetById implements properties.ServiceInterface
func (service *propertyService) GetById(id int) (data property.Core, err error) {
	data, err = service.propertyRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return property.Core{}, helper.ServiceErrorMsg(err)
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
