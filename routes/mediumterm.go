package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteMediumterm(app *echo.Echo) {
	app.GET("/mediumterm", controller.PrintMediumterm, middleware.JWTAuthMiddleware)
	app.GET("/mediumterm/:id", controller.PrintMediumtermByID, middleware.JWTAuthMiddleware)
	app.POST("/mediumterm", controller.PrintCreateMediumterm, middleware.JWTAuthMiddleware)
	app.PUT("/mediumterm/:id", controller.PrintUpdateMediumterm, middleware.JWTAuthMiddleware)
	app.DELETE("/mediumterm/:id", controller.PrintDeleteMediumterm, middleware.JWTAuthMiddleware)
}