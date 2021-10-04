package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintLongtermbyage(c echo.Context) error {
	term := database.GetLongtermbyage()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Longtermbyage",
		"data":    term,
	})
}

func PrintCreateLongtermbyage(c echo.Context) error {
	var newTerm model.Longtermbyage
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longtermbyage",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateLongtermbyage(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longtermbyage",
		"data":    newTerm,
	})
}

func PrintUpdateLongtermbyage(c echo.Context) error {
	id := c.Param("id")

	var term model.Longtermbyage
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longtermbyage",
			"error":   err.Error(),
		})
	}
	database.UpdateLongtermbyage(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longtermbyage",
		"data":    term,
	})
}

func PrintDeleteLongtermbyage(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLongtermbyage(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Longtermbyage",
	})
}