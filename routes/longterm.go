package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteLongterm(app *echo.Echo) {
	app.GET("/longterm", controller.PrintLongterm, middleware.JWTAuthMiddleware)
	app.GET("/longterm/:id", controller.PrintLongtermByID, middleware.JWTAuthMiddleware)
	app.POST("/longterm", controller.PrintCreateLongterm, middleware.JWTAuthMiddleware)
	app.PUT("/longterm/:id", controller.PrintUpdateLongterm, middleware.JWTAuthMiddleware)
	app.DELETE("/longterm/:id", controller.PrintDeleteLongterm, middleware.JWTAuthMiddleware)

}
