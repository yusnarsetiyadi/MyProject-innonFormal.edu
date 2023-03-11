package main

import (
	"fmt"
	"myproject/innonformaledu/internal/config"
	"myproject/innonformaledu/internal/factory"
	"myproject/innonformaledu/internal/middlewares"

	"myproject/innonformaledu/utils/database/mysql"
	// "myproject/innonformaledu/utils/database/postgresql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// config and connection db
	cfg := config.GetConfig()
	db := mysql.InitMySQLDB(cfg)
	// db := postgresql.InitPostgreSQLDB(cfg)

	// echo framework
	e := echo.New()

	// features
	factory.InitFactory(e, db)

	// middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
