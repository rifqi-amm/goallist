package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteUser(app *echo.Echo) {
	app.GET("/user", controller.PrintUser, middleware.JWTAuthMiddleware)
	app.GET("/user/:id", controller.PrintUserByID, middleware.JWTAuthMiddleware)
	app.POST("/user", controller.PrintCreateUser, middleware.JWTAuthMiddleware)
	app.PUT("/user/:id", controller.PrintUpdateUser, middleware.JWTAuthMiddleware)
	app.DELETE("/user/:id", controller.PrintDeleteUser, middleware.JWTAuthMiddleware)
}