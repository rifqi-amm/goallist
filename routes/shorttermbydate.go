package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteShorttermbydate(app *echo.Echo) {
	app.GET("/shorttermbydate", controller.PrintShorttermbydate, middleware.JWTAuthMiddleware)
	app.GET("/shorttermbydate/:id", controller.PrintShorttermbydateByID, middleware.JWTAuthMiddleware)
	app.POST("/shorttermbydate", controller.PrintCreateShorttermbydate, middleware.JWTAuthMiddleware)
	app.PUT("/shorttermbydate/:id", controller.PrintUpdateShorttermbydate, middleware.JWTAuthMiddleware)
	app.DELETE("/shorttermbydate/:id", controller.PrintDeleteShorttermbydate, middleware.JWTAuthMiddleware)
}