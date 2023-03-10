package delivery

import (
	"myproject/innonformaledu/internal/features/auth"
	"myproject/innonformaledu/utils/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

/*

	Handler Input Output Application

*/

type AuthHandler struct {
	authService auth.ServiceInterface
}

func (handler *AuthHandler) Login(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to bind data."))
	}

	dataCore := ToCore(userInput)
	result, token, err := handler.authService.Login(dataCore)

	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to Login. "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Success.", FromCore(result, token)))
}
