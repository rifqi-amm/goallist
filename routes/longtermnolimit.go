package routes

import (
	"goallist/controller"
	"goallist/middleware"

	"github.com/labstack/echo/v4"
)

func RouteLongtermnolimit(app *echo.Echo) {
	app.GET("/longtermnolimit", controller.PrintLongtermnolimit, middleware.JWTAuthMiddleware)
	app.GET("/longtermnolimit/:id", controller.PrintLongtermnolimitByID, middleware.JWTAuthMiddleware)
	app.POST("/longtermnolimit", controller.PrintCreateLongtermnolimit, middleware.JWTAuthMiddleware)
	app.PUT("/longtermnolimit/:id", controller.PrintUpdateLongtermnolimit, middleware.JWTAuthMiddleware)
	app.DELETE("/longtermnolimit/:id", controller.PrintDeleteLongtermnolimit, middleware.JWTAuthMiddleware)
}
