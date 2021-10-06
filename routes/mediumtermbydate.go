package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteMediumtermbydate(app *echo.Echo) {
	app.GET("/mediumtermbydate", controller.PrintMediumtermbydate, middleware.JWTAuthMiddleware)
	app.GET("/mediumtermbydate/:id", controller.PrintMediumtermbydateByID, middleware.JWTAuthMiddleware)
	app.POST("/mediumtermbydate", controller.PrintCreateMediumtermbydate, middleware.JWTAuthMiddleware)
	app.PUT("/mediumtermbydate/:id", controller.PrintUpdateMediumtermbydate, middleware.JWTAuthMiddleware)
	app.DELETE("/mediumtermbydate/:id", controller.PrintDeleteMediumtermbydate, middleware.JWTAuthMiddleware)
}