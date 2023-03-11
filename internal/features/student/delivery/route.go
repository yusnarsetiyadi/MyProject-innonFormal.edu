package delivery

import (
	"myproject/innonformaledu/internal/features/student"
	"myproject/innonformaledu/internal/middlewares"

	"github.com/labstack/echo/v4"
)

/*

	Route Endpoint

*/

func New(service student.ServiceInterface, e *echo.Echo) {
	group := *echo.New().Group("/student")
	handler := &StudentDelivery{
		studentService: service,
	}

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetById)
	group.POST("", handler.Create)
	group.PUT("/:id", handler.Update, middlewares.JWTMiddleware())
	group.DELETE("/:id", handler.Delete, middlewares.JWTMiddleware())
}
