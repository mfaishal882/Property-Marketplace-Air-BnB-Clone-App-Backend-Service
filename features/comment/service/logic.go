package service

import (
	"api-airbnb-alta/features/comment"
	"api-airbnb-alta/utils/helper"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type commentService struct {
	comentRepository comment.RepositoryInterface
	validate         *validator.Validate
}

func New(repo comment.RepositoryInterface) comment.ServiceInterface {
	return &commentService{
		comentRepository: repo,
		validate:         validator.New(),
	}
}

func (service *commentService) Create(input comment.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.comentRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *commentService) GetAll(query string) (data []comment.Core, err error) {

	data, err = service.comentRepository.GetAll()

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

func (service *commentService) GetById(id int) (data comment.Core, err error) {
	data, err = service.comentRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return comment.Core{}, helper.ServiceErrorMsg(err)
	}

	if data == (comment.Core{}) {
		helper.LogDebug("Get data success. No data.")
		return comment.Core{}, errors.New("Get data success. No data.")
	}

	return data, err

}

func (service *commentService) Update(input comment.Core, id int, c echo.Context) error {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.comentRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.comentRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *commentService) Delete(id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.comentRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.comentRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
