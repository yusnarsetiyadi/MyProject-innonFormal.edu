package delivery

import (
	"myproject/innonformaledu/internal/features/joinclass"
	"myproject/innonformaledu/internal/middlewares"

	"github.com/labstack/echo/v4"
)

/*

	Route Endpoint

*/

func New(service joinclass.ServiceInterface, e *echo.Echo) {
	group := *echo.New().Group("/join-class")
	handler := &JoinClassDelivery{
		joinclassService: service,
	}

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetById)
	group.POST("", handler.Create)
	group.PUT("/:id", handler.Update, middlewares.JWTMiddleware())
	group.DELETE("/:id", handler.Delete, middlewares.JWTMiddleware())
}
