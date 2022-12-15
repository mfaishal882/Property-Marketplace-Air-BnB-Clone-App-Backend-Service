package delivery

import (
	property "api-airbnb-alta/features/property"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type PropertyDelivery struct {
	propertyService property.ServiceInterface
}

func New(service property.ServiceInterface, e *echo.Echo) {
	handler := &PropertyDelivery{
		propertyService: service,
	}

	e.GET("/properties", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/properties/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/properties", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/properties/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/properties/:id", handler.Delete, middlewares.JWTMiddleware())
	e.GET("/properties/:id/images", handler.GetPropertyImages, middlewares.JWTMiddleware())
	e.GET("/properties/:id/comments", handler.GetPropertyComments, middlewares.JWTMiddleware())
	e.GET("/properties/:id/availability", handler.GetAvailability, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja
}

func (delivery *PropertyDelivery) GetAll(c echo.Context) error {
	queryPropertyName := c.QueryParam("property_name")
	queryCity := c.QueryParam("city")
	queryPropertyType := c.QueryParam("property_type")

	helper.LogDebug("\n\n\nULALA")

	// debug cek query param masuk
	helper.LogDebug("\n isi queryPropertyName = ", queryPropertyName)
	helper.LogDebug("\n isi queryCity= ", queryCity)
	helper.LogDebug("\n isi queryPropertyType = ", queryPropertyType)

	results, err := delivery.propertyService.GetAll(queryPropertyName, queryCity, queryPropertyType)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *PropertyDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.propertyService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreById(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *PropertyDelivery) Create(c echo.Context) error {
	propertyInput := InsertRequest{}
	errBind := c.Bind(&propertyInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(propertyInput)
	helper.LogDebug("\n dataCore = ", dataCore)
	err := delivery.propertyService.Create(dataCore, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *PropertyDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	propertyInput := InsertRequest{}
	errBind := c.Bind(&propertyInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	propertyData, errGet := delivery.propertyService.GetById(id)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	}

	fmt.Println("\n\nProp data = ", propertyData)
	fmt.Println("\n\nid = ", userId, " = prop user id =", propertyData.UserID)

	if userId != int(propertyData.UserID) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	}

	dataCore := toCore(propertyInput)
	err := delivery.propertyService.Update(dataCore, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *PropertyDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	propertyData, errGet := delivery.propertyService.GetById(id)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	}

	fmt.Println("\n\nProp data = ", propertyData)
	fmt.Println("\n\nid = ", userId, " = prop user id =", propertyData.UserID)

	if userId != int(propertyData.UserID) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	}

	err := delivery.propertyService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}

func (delivery *PropertyDelivery) GetPropertyImages(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.propertyService.GetPropertyImages(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromPropertyImagesList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *PropertyDelivery) GetPropertyComments(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.propertyService.GetPropertyComments(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromPropertyCommentList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *PropertyDelivery) GetAvailability(c echo.Context) error {
	idParam := c.Param("id")
	checkinDate := c.QueryParam("checkin_date")
	checkoutDate := c.QueryParam("checkout_date")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	var layout2 = "2006-01-02"
	checkinDateTime, errParse1 := time.Parse(layout2, checkinDate)
	checkoutDateTime, errParse2 := time.Parse(layout2, checkoutDate)

	if errParse1 != nil || errParse2 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Parse "+errParse1.Error()+" dan "+errParse2.Error()))
	}

	results, err := delivery.propertyService.GetAvailbility(uint(id), checkinDateTime, checkoutDateTime)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	propertyData, errGetData := delivery.propertyService.GetById(id)

	fmt.Println("\n\n\n\n Hasil Checkin date", checkinDateTime)
	fmt.Println("\nHasil Checkout date", checkoutDateTime)
	fmt.Println("\nHasil result avail", results)
	fmt.Println("\nGet property by id", propertyData)

	if errGetData != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreCheckAvailbility(propertyData, checkinDate, checkoutDate, results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}
