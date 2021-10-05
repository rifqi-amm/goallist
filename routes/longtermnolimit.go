package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteLongtermnolimit(app *echo.Echo) {
	app.GET("/longtermnolimit", controller.PrintLongtermnolimit)
	app.POST("/longtermnolimit", controller.PrintCreateLongtermnolimit)
	app.PUT("/longtermnolimit/:id", controller.PrintUpdateLongtermnolimit)
	app.DELETE("/longtermnolimit/:id", controller.PrintDeleteLongtermnolimit)
}
