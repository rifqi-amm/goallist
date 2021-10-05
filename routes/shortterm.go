package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteShortterm(app *echo.Echo) {
	app.GET("/shortterm", controller.PrintShortterm)
	app.POST("/shortterm", controller.PrintCreateShortterm)
	app.PUT("/shortterm/:id", controller.PrintUpdateShortterm)
	app.DELETE("/shortterm/:id", controller.PrintDeleteShortterm)
}