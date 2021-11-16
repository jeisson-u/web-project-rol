package apiController

import (
	"net/http"

	apimodels "github.com/jeisson-u/web-project-rol.git/apiModels"
	"github.com/jeisson-u/web-project-rol.git/database"
	"github.com/labstack/echo/v4"
)

type SecurityController struct {
}

func (SecurityController) Login(c echo.Context) (err error) {
	modelIn := new(apimodels.LoginModelIn)

	if err = c.Bind(modelIn); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userDb := database.UserDatabase{}
	userDb.Start()
	defer userDb.Close()

	user, err := userDb.GetByEmailAndPassword(modelIn.Email, modelIn.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (SecurityController) RefreshToken(c echo.Context) error {
	return c.String(http.StatusOK, "1")
}
