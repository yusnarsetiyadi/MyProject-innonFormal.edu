package delivery

import (
	"myproject/innonformaledu/internal/features/class"
	"myproject/innonformaledu/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

/*

	Handler Input Output Application

*/

type ClassDelivery struct {
	classService class.ServiceInterface
}

func (delivery *ClassDelivery) GetAll(c echo.Context) error {
	results, err := delivery.classService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *ClassDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.classService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *ClassDelivery) Create(c echo.Context) error {
	input := ClassRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	dataCore := toCore(input)
	err := delivery.classService.Create(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ClassDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	input := ClassRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}
	dataCore := toCore(input)
	err := delivery.classService.Update(dataCore, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

func (delivery *ClassDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.classService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
