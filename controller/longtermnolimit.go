package controller

import (
	"net/http"

	"goallist/database"
	"goallist/model"

	"github.com/labstack/echo/v4"
)

func PrintLongtermnolimit(c echo.Context) error {
	term := database.GetLongtermnolimit()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print Longtermnolimit",
		"data":    term,
	})
}

func PrintLongtermnolimitByID(c echo.Context) error {
	id := c.Param("id")
	term := database.GetLongtermnolimitByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to show a goal list",
		"data":    term,
	})
}

func PrintCreateLongtermnolimit(c echo.Context) error {
	var newTerm model.Longtermnolimit
	if err := c.Bind(&newTerm); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longtermnolimit",
			"error":   err.Error(),
		})
	}

	newTerm = database.CreateLongtermnolimit(newTerm)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longtermnolimit",
		"data":    newTerm,
	})
}

func PrintUpdateLongtermnolimit(c echo.Context) error {
	id := c.Param("id")

	var term model.Longtermnolimit
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Succes to crate data into Longtermnolimit",
			"error":   err.Error(),
		})
	}
	database.UpdateLongtermnolimit(id, term)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to print data Longtermnolimit",
		"data":    term,
	})
}

func PrintDeleteLongtermnolimit(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLongtermnolimit(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Succes to delete data from Longtermnolimit",
	})
}