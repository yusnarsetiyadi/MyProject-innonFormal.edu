package delivery

import (
	"myproject/innonformaledu/internal/features/auth"

	"github.com/labstack/echo/v4"
)

/*

	Route Endpoint

*/

func New(service auth.ServiceInterface, e *echo.Echo) {
	group := *echo.New().Group("/auth")
	handler := &AuthHandler{
		authService: service,
	}

	group.POST("/login", handler.Login)
}
