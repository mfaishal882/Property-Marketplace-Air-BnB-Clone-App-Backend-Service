package delivery

import (
	"api-airbnb-alta/features/booking"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type BookingDelivery struct {
	bookingService booking.ServiceInterface
}

func New(service booking.ServiceInterface, e *echo.Echo) {
	handler := &BookingDelivery{
		bookingService: service,
	}

	e.GET("/bookings", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/bookings/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/bookings", handler.Create, middlewares.JWTMiddleware())
}

func (delivery *BookingDelivery) GetAll(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, errors.New("Failed load user id from JWT token, please check again."))
	}

	results, err := delivery.bookingService.GetAll(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all users", dataResponse))
}

func (delivery *BookingDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, errors.New("Failed load user id from JWT token, please check again."))
	}
	results, err := delivery.bookingService.GetById(id, userId)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *BookingDelivery) Create(c echo.Context) error {
	bookingInput := InsertRequest{}
	errBind := c.Bind(&bookingInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(bookingInput)
	helper.LogDebug("\n dataCore = ", dataCore)
	err := delivery.bookingService.Create(dataCore, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}
