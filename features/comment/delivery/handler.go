package delivery

import (
	"api-airbnb-alta/features/comment"
	"api-airbnb-alta/middlewares"
	"api-airbnb-alta/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CommentDelivery struct {
	commentService comment.ServiceInterface
}

func New(service comment.ServiceInterface, e *echo.Echo) {
	handler := &CommentDelivery{
		commentService: service,
	}

	e.GET("/comments", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/comments/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/comments", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/comments/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/comments/:id", handler.Delete, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *CommentDelivery) GetAll(c echo.Context) error {
	query := c.QueryParam("title")
	helper.LogDebug("isi query = ", query)
	results, err := delivery.commentService.GetAll(query)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all users", dataResponse))
}

func (delivery *CommentDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.commentService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *CommentDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	userId := middlewares.ExtractTokenUserId(c)
	dataCore := toCore(userInput, uint(userId))
	err := delivery.commentService.Create(dataCore, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *CommentDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userInput := UpdateRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	userId := middlewares.ExtractTokenUserId(c)
	dataCore := toCore(userInput, uint(userId))
	err := delivery.commentService.Update(dataCore, id, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *CommentDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.commentService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
