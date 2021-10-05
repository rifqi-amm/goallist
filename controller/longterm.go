package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintLongterm(c echo.Context) error {
	term := database.GetLongterm()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Longterm",
		"data":    term,
	})
}

func PrintLongtermByID(c echo.Context) error {
	id := c.Param("id")
	term := database.GetLongtermByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to show a goal list",
		"data":    term,
	})
}

func PrintCreateLongterm(c echo.Context) error {
	var newTerm model.Longterm
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longterm",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateLongterm(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longterm",
		"data":    newTerm,
	})
}

func PrintUpdateLongterm(c echo.Context) error {
	id := c.Param("id")

	var term model.Longterm
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longterm",
			"error":   err.Error(),
		})
	}
	database.UpdateLongterm(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longterm",
		"data":    term,
	})
}

func PrintDeleteLongterm(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLongterm(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Longterm",
	})
}