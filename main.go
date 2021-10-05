package main

import (
	"goallist/config"
	"goallist/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	config.InitMigration()

	app := echo.New()

	
	routes.RouteLoginuser(app)
	routes.RouteAllterm(app)
	routes.RouteShortterm(app)
	routes.RouteShorttermbydate(app)
	routes.RouteMediumterm(app)
	routes.RouteMediumtermbydate(app)
	routes.RouteLongterm(app)
	routes.RouteLongtermbyage(app)
	routes.RouteLongtermnolimit(app)
	app.Start(":8080")
}