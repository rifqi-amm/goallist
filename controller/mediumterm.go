package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintMediumterm(c echo.Context) error {
	term := database.GetMediumterm()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Mediumterm",
		"data":    term,
	})
}

func PrintMediumtermByID(c echo.Context) error {
	id := c.Param("id")
	term := database.GetMediumtermByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to show a goal list",
		"data":    term,
	})
}

func PrintCreateMediumterm(c echo.Context) error {
	var newTerm model.Mediumterm
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Mediumterm",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateMediumterm(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Mediumterm",
		"data":    newTerm,
	})
}

func PrintUpdateMediumterm(c echo.Context) error {
	id := c.Param("id")

	var term model.Mediumterm
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Mediumterm",
			"error":   err.Error(),
		})
	}
	database.UpdateMediumterm(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Mediumterm",
		"data":    term,
	})
}

func PrintDeleteMediumterm(c echo.Context) error {
	id := c.Param("id")
	database.DeleteMediumterm(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Mediumterm",
	})
}