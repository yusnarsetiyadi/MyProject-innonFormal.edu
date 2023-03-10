package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] status=${status} method=${method} host=${host} path=${path} latency_human=${latency_human} id=${id} remote_ip=${remote_ip} uri=${uri} error=${error}` + "\n=========================================\n",
	}))
}
