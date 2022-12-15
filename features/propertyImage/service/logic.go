package service

import (
	"api-airbnb-alta/features/propertyImage"
	"api-airbnb-alta/utils/helper"
	"api-airbnb-alta/utils/thirdparty"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type propertyImageService struct {
	propertyImageRepository propertyImage.RepositoryInterface
	validate                *validator.Validate
}

func New(repo propertyImage.RepositoryInterface) propertyImage.ServiceInterface {
	return &propertyImageService{
		propertyImageRepository: repo,
		validate:                validator.New(),
	}
}

func (service *propertyImageService) Create(input propertyImage.Core, c echo.Context) (err error) {

	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.Upload(c)
		if err != nil {
			helper.LogDebug("Failed Upload : ", res, " Error ", err)
			return errors.New("Failed. Cannot Upload Data.")
		}
		helper.LogDebug("Success Upload : ", res)
		input.ImageUrl = res
	} else {
		helper.LogDebug("Failed Upload : using default picture")
		input.ImageUrl = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}

	errCreate := service.propertyImageRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *propertyImageService) GetAll(query string) (data []propertyImage.Core, err error) {

	data, err = service.propertyImageRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

func (service *propertyImageService) GetById(id int) (data propertyImage.Core, err error) {
	data, err = service.propertyImageRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return propertyImage.Core{}, helper.ServiceErrorMsg(err)
	}

	if data == (propertyImage.Core{}) {
		helper.LogDebug("Get data success. No data.")
		return propertyImage.Core{}, errors.New("Get data success. No data.")
	}

	return data, err

}

func (service *propertyImageService) Update(input propertyImage.Core, id int, c echo.Context) error {

	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.propertyImageRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.Upload(c)
		if err != nil {
			helper.LogDebug("Failed Upload : ", res, " Error ", err)
			return errors.New("Failed. Cannot Upload Data.")
		}
		helper.LogDebug("Success Upload : ", res)
		input.ImageUrl = res
	} else {
		helper.LogDebug("Failed Upload : using default picture")
		input.ImageUrl = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}

	// proses
	err := service.propertyImageRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *propertyImageService) Delete(id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.propertyImageRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.propertyImageRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

// GetPropertyById implements propertyImage.ServiceInterface
func (service *propertyImageService) GetPropertyById(id int) (data propertyImage.Property, err error) {
	data, err = service.propertyImageRepository.GetPropertyById(id)
	if err != nil {
		log.Error(err.Error())
		return propertyImage.Property{}, helper.ServiceErrorMsg(err)
	}

	if data == (propertyImage.Property{}) {
		helper.LogDebug("Get data success. No data.")
		return propertyImage.Property{}, errors.New("Get data success. No data.")
	}

	return data, err
}
