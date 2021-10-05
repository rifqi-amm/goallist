package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteShorttermbydate(app *echo.Echo) {
	app.GET("/shorttermbydate", controller.PrintShorttermbydate)
	app.POST("/shorttermbydate", controller.PrintCreateShorttermbydate)
	app.PUT("/shorttermbydate/:id", controller.PrintUpdateShorttermbydate)
	app.DELETE("/shorttermbydate/:id", controller.PrintDeleteShorttermbydate)
}