package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteMediumtermbydate(app *echo.Echo) {
	app.GET("/mediumtermbydate", controller.PrintMediumtermbydate)
	app.POST("/mediumtermbydate", controller.PrintCreateMediumtermbydate)
	app.PUT("/mediumtermbydate/:id", controller.PrintUpdateMediumtermbydate)
	app.DELETE("/mediumtermbydate/:id", controller.PrintDeleteMediumtermbydate)
}