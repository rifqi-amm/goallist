package controller

import (
	"goallist/database"
	"goallist/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PrintUser(c echo.Context) error {
	datauser := database.GetUser()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print User",
		"data":    datauser,
	})
}

func PrintUserByID(c echo.Context) error {
	id := c.Param("id")
	datauser := database.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to show a goal list",
		"data":    datauser,
	})
}


func PrintCreateUser(c echo.Context) error {
	var userdataNew model.User
	if err := c.Bind(&userdataNew); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data User",
			"error":   err.Error(),
		})
	}

	userdataNew = database.CreateUser(userdataNew)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data User",
		"data":    userdataNew,
	})
}

func PrintUpdateUser(c echo.Context) error {
	id := c.Param("id")

	var datauser model.User
	if err := c.Bind(&datauser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data user",
			"error":   err.Error(),
		})
	}
	database.UpdateUser(id, datauser)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data User",
		"data":    datauser,
	})
}

func PrintDeleteUser(c echo.Context) error {
	id := c.Param("id")
	database.DeleteUser(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from User",
	})
}