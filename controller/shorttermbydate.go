package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintShorttermbydate(c echo.Context) error {
	term := database.GetShorttermbydate()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Shorttermbydate",
		"data":    term,
	})
}

func PrintCreateShorttermbydate(c echo.Context) error {
	var newTerm model.Shorttermbydate
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Shorttermbydate",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateShorttermbydate(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Shorttermbydate",
		"data":    newTerm,
	})
}

func PrintUpdateShorttermbydate(c echo.Context) error {
	id := c.Param("id")

	var term model.Shorttermbydate
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Shorttermbydate",
			"error":   err.Error(),
		})
	}
	database.UpdateShorttermbydate(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Shorttermbydate",
		"data":    term,
	})
}

func PrintDeleteShorttermbydate(c echo.Context) error {
	id := c.Param("id")
	database.DeleteShorttermbydate(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Shorttermbydate",
	})
}