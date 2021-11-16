package webController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
}

func (HomeController) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{})
}
