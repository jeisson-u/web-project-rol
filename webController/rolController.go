package webController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RolController struct {
}

func (RolController) RolsPage(c echo.Context) error {
	return c.Render(http.StatusOK, "rols", echo.Map{})
}
