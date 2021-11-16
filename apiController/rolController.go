package apiController

import (
	"net/http"

	"github.com/jeisson-u/web-project-rol.git/database"
	"github.com/labstack/echo/v4"
)

type RolController struct {
}

func (RolController) GetAll(c echo.Context) (err error) {

	userDb := database.RolDatabase{}
	userDb.Start()
	defer userDb.Close()

	data, err := userDb.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
