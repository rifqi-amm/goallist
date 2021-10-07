package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteUser(app *echo.Echo) {
	app.GET("/User", controller.PrintUser, middleware.JWTAuthMiddleware)
	app.GET("/User/:id", controller.PrintUserByID, middleware.JWTAuthMiddleware)
	app.POST("/User", controller.PrintCreateUser, middleware.JWTAuthMiddleware)
	app.PUT("/User/:id", controller.PrintUpdateUser, middleware.JWTAuthMiddleware)
	app.DELETE("/User/:id", controller.PrintDeleteUser, middleware.JWTAuthMiddleware)
}