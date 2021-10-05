package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteMediumterm(app *echo.Echo) {
	app.GET("/mediumterm", controller.PrintMediumterm)
	app.GET("/mediumterm/:id", controller.PrintMediumtermByID)
	app.POST("/mediumterm", controller.PrintCreateMediumterm)
	app.PUT("/mediumterm/:id", controller.PrintUpdateMediumterm)
	app.DELETE("/mediumterm/:id", controller.PrintDeleteMediumterm)
}