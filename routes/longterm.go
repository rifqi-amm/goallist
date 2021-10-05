package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func RouteLongterm(app *echo.Echo) {
	app.GET("/longterm", controller.PrintLongterm)
	app.GET("/longterm/:id", controller.PrintLongtermByID)
	app.POST("/longterm", controller.PrintCreateLongterm)
	app.PUT("/longterm/:id", controller.PrintUpdateLongterm)
	app.DELETE("/longterm/:id", controller.PrintDeleteLongterm)

}
