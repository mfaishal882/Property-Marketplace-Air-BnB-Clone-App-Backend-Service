package delivery

import (
	property "api-airbnb-alta/features/property"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"net/http"
	"strconv"

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
	err := delivery.propertyService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
