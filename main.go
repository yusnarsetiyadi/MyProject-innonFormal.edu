package main

import (
	"fmt"
	"myproject/innonformaledu/config"
	"myproject/innonformaledu/factory"
	"myproject/innonformaledu/middlewares"

	"myproject/innonformaledu/utils/database/mysql"
	// "myproject/innonformaledu/utils/database/postgresql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMySQLDB(cfg)
	// db := postgresql.InitPostgreSQLDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	//middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
