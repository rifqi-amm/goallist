package main

import (
	"goallist/config"
	"goallist/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	config.InitMigration()

	app := echo.New()
	routes.PritntData(app)
	app.Start(":8080")
}