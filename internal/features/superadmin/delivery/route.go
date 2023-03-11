package delivery

import (
	"myproject/innonformaledu/internal/features/superadmin"
	"myproject/innonformaledu/internal/middlewares"

	"github.com/labstack/echo/v4"
)

/*

	Route Endpoint

*/

func New(service superadmin.ServiceInterface, e *echo.Echo) {
	group := *echo.New().Group("/super-admin")
	handler := &SuperAdminDelivery{
		superadminService: service,
	}

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetById)
	group.POST("", handler.Create)
	group.PUT("/:id", handler.Update, middlewares.JWTMiddleware())
	group.DELETE("/:id", handler.Delete, middlewares.JWTMiddleware())
}
