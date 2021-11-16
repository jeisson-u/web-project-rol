package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ClientController struct {
}

func (ClientController) CreateClient(c echo.Context) error {
	return c.String(http.StatusOK, "1")
}
