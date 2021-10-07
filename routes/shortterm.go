package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteShortterm(app *echo.Echo) {
	app.GET("/shortterm", controller.PrintShortterm, middleware.JWTAuthMiddleware)
	app.GET("/shortterm/:id", controller.PrintShorttermByID, middleware.JWTAuthMiddleware)
	app.POST("/shortterm", controller.PrintCreateShortterm, middleware.JWTAuthMiddleware)
	app.PUT("/shortterm/:id", controller.PrintUpdateShortterm, middleware.JWTAuthMiddleware)
	app.DELETE("/shortterm/:id", controller.PrintDeleteShortterm, middleware.JWTAuthMiddleware)
}