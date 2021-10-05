package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteAllterm(app *echo.Echo) {
	app.GET("/allterm", controller.PrintAllterm)
}