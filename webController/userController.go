package webController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func (UserController) UsersPage(c echo.Context) error {
	return c.Render(http.StatusOK, "users", echo.Map{})
}
