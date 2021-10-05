package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintShortterm(c echo.Context) error {
	term := database.GetShortterm()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Shortterm",
		"data":    term,
	})
}

func PrintShorttermByID(c echo.Context) error {
	id := c.Param("id")
	term := database.GetShorttermByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to show a goal list",
		"data":    term,
	})
}

func PrintCreateShortterm(c echo.Context) error {
	var newTerm model.Shortterm
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Shortterm",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateShortterm(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Shortterm",
		"data":    newTerm,
	})
}

func PrintUpdateShortterm(c echo.Context) error {
	id := c.Param("id")

	var term model.Shortterm
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Shortterm",
			"error":   err.Error(),
		})
	}
	database.UpdateShortterm(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Shortterm",
		"data":    term,
	})
}

func PrintDeleteShortterm(c echo.Context) error {
	id := c.Param("id")
	database.DeleteShortterm(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from shortterm",
	})
}