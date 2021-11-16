package webController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SecurityController struct {
}

func (SecurityController) LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login", echo.Map{})
}
