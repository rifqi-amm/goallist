package routes

import (
	"goallist/controller"

	"github.com/labstack/echo/v4"
)

func PritntData(app *echo.Echo) {
	// Print All Goal List
	app.GET("/allterm", controller.PrintAllterm)

	// Print A Term
	app.GET("/shortterm", controller.PrintShortterm)
	app.GET("/shorttermbydate", controller.PrintShorttermbydate)
	app.GET("/mediumterm", controller.PrintMediumterm)
	app.GET("/mediumtermbydate", controller.PrintMediumtermbydate)
	app.GET("/longterm", controller.PrintLongterm)
	app.GET("/longtermbyage", controller.PrintLongtermbyage)
	app.GET("/longtermnolimit", controller.PrintLongtermnolimit)

	// Print A Create
	app.POST("/shortterm", controller.PrintCreateShortterm)
	app.POST("/shorttermbydate", controller.PrintCreateShorttermbydate)
	app.POST("/mediumterm", controller.PrintCreateMediumterm)
	app.POST("/mediumtermbydate", controller.PrintCreateMediumtermbydate)
	app.POST("/longterm", controller.PrintCreateLongterm)
	app.POST("/longtermbyage", controller.PrintCreateLongtermbyage)
	app.POST("/longtermnolimit", controller.PrintCreateLongtermnolimit)

	// Print A Update
	app.PUT("/shortterm/:id", controller.PrintUpdateShortterm)
	app.PUT("/shorttermbydate/:id", controller.PrintUpdateShorttermbydate)
	app.PUT("/mediumterm/:id", controller.PrintUpdateMediumterm)
	app.PUT("/mediumtermbydate/:id", controller.PrintUpdateMediumtermbydate)
	app.PUT("/longterm/:id", controller.PrintUpdateLongterm)
	app.PUT("/longtermbyage/:id", controller.PrintUpdateLongtermbyage)
	app.PUT("/longtermnolimit/:id", controller.PrintUpdateLongtermnolimit)

	// Print A Delete
	app.DELETE("/shortterm/:id", controller.PrintDeleteShortterm)
	app.DELETE("/shorttermbydate/:id", controller.PrintDeleteShorttermbydate)
	app.DELETE("/mediumterm/:id", controller.PrintDeleteMediumterm)
	app.DELETE("/mediumtermbydate/:id", controller.PrintDeleteMediumtermbydate)
	app.DELETE("/longterm/:id", controller.PrintDeleteLongterm)
	app.DELETE("/longtermbyage/:id", controller.PrintDeleteLongtermbyage)
	app.DELETE("/longtermnolimit/:id", controller.PrintDeleteLongtermnolimit)
}