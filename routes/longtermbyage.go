package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteLongtermbyage(app *echo.Echo) {
	app.GET("/longtermbyage", controller.PrintLongtermbyage)
	app.GET("/longtermbyage/:id", controller.PrintLongtermbyageByID)
	app.POST("/longtermbyage", controller.PrintCreateLongtermbyage)
	app.PUT("/longtermbyage/:id", controller.PrintUpdateLongtermbyage)
	app.DELETE("/longtermbyage/:id", controller.PrintDeleteLongtermbyage)
}