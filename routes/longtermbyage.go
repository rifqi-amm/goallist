package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteLongtermbyage(app *echo.Echo) {
	app.GET("/longtermbyage", controller.PrintLongtermbyage, middleware.JWTAuthMiddleware)
	app.GET("/longtermbyage/:id", controller.PrintLongtermbyageByID, middleware.JWTAuthMiddleware)
	app.POST("/longtermbyage", controller.PrintCreateLongtermbyage, middleware.JWTAuthMiddleware)
	app.PUT("/longtermbyage/:id", controller.PrintUpdateLongtermbyage, middleware.JWTAuthMiddleware)
	app.DELETE("/longtermbyage/:id", controller.PrintDeleteLongtermbyage, middleware.JWTAuthMiddleware)
}