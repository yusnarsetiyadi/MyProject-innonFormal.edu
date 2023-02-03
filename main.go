package main

import (
	"fmt"
	"myproject/innonformaledu/config"
	"myproject/innonformaledu/middlewares"

	"myproject/innonformaledu/utils/database/mysql"
	// "myproject/innonformaledu/utils/database/postgresql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	mysql.InitMySQLDB(cfg)
	// postgresql.InitPostgreSQLDB(cfg)

	e := echo.New()

	//middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
