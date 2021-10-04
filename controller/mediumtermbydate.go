package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintMediumtermbydate(c echo.Context) error {
	term := database.GetMediumtermbydate()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Mediumtermbydate",
		"data":    term,
	})
}

func PrintCreateMediumtermbydate(c echo.Context) error {
	var newTerm model.Mediumtermbydate
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Mediumtermbydate",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateMediumtermbydate(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Mediumtermbydate",
		"data":    newTerm,
	})
}

func PrintUpdateMediumtermbydate(c echo.Context) error {
	id := c.Param("id")

	var term model.Mediumtermbydate
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Mediumtermbydate",
			"error":   err.Error(),
		})
	}
	database.UpdateMediumtermbydate(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Mediumtermbydate",
		"data":    term,
	})
}

func PrintDeleteMediumtermbydate(c echo.Context) error {
	id := c.Param("id")
	database.DeleteMediumtermbydate(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Mediumtermbydate",
	})
}