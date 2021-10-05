package controller

import (
	"goallist/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PrintAllterm(c echo.Context) error {
	termShortterm := database.GetShortterm()
	termShorttermbydate := database.GetShorttermbydate()
	termMediumterm := database.GetMediumterm()
	termMediumtermbydate := database.GetMediumtermbydate()
	termLongterm := database.GetLongterm()
	termLongtermbyage := database.GetLongtermbyage()
	termLongtermnolimit := database.GetLongtermnolimit()
	Data := echo.Map{
		"Short Term" : termShortterm,
		"Short Term by Date" : termShorttermbydate,
		"Medium Term" : termMediumterm,
		"Medium Term by Date" : termMediumtermbydate,
		"Long Term" : termLongterm,
		"Long Term by Age" : termLongtermbyage,
		"Long Term No Limit" : termLongtermnolimit,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "succes to print all term",
		"data":    Data,
	})
}


