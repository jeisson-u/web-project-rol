package webController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BillingController struct {
}

func (BillingController) BillingPage(c echo.Context) error {
	return c.Render(http.StatusOK, "billing", echo.Map{})
}
