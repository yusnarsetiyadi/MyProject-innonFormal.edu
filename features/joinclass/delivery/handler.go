package delivery

import (
	"myproject/innonformaledu/features/joinclass"
	"myproject/innonformaledu/middlewares"
	"myproject/innonformaledu/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type JoinClassDelivery struct {
	joinclassService joinclass.ServiceInterface
}

func New(service joinclass.ServiceInterface, e *echo.Echo) {
	handler := &JoinClassDelivery{
		joinclassService: service,
	}

	e.GET("/join-classes", handler.GetAll)
	e.GET("/join-classes/:id", handler.GetById)
	e.POST("/join-classes", handler.Create)
	e.PUT("/join-classes/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/join-classes/:id", handler.Delete, middlewares.JWTMiddleware())
}

func (delivery *JoinClassDelivery) GetAll(c echo.Context) error {
	results, err := delivery.joinclassService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *JoinClassDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.joinclassService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *JoinClassDelivery) Create(c echo.Context) error {
	input := JoinClassRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	dataCore := toCore(input)
	err := delivery.joinclassService.Create(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *JoinClassDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	input := JoinClassRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	dataCore := toCore(input)
	err := delivery.joinclassService.Update(dataCore, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

func (delivery *JoinClassDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.joinclassService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
