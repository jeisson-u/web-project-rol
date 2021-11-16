package apiController

import (
	"net/http"
	"strconv"

	apimodels "github.com/jeisson-u/web-project-rol.git/apiModels"
	"github.com/jeisson-u/web-project-rol.git/database"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func (UserController) CreateUser(c echo.Context) (err error) {

	modelIn := new(apimodels.CreateUserModelIn)

	if err = c.Bind(modelIn); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userDb := database.UserDatabase{}
	userDb.Start()
	defer userDb.Close()

	err = userDb.Create(modelIn.Name, modelIn.Email, modelIn.Password, 1)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
func (UserController) UpdateUser(c echo.Context) (err error) {
	modelIn := new(apimodels.UpdateUserModelIn)

	if err = c.Bind(modelIn); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userDb := database.UserDatabase{}
	userDb.Start()
	defer userDb.Close()

	err = userDb.Update(modelIn.Id, modelIn.Name, modelIn.Password, modelIn.State)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "")
}
func (UserController) RemoveUser(c echo.Context) (err error) {

	userId := c.Param("userId")

	userDb := database.UserDatabase{}
	userDb.Start()
	defer userDb.Close()

	id, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = userDb.Delete(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "")
}

func (UserController) GetAll(c echo.Context) (err error) {

	userDb := database.UserDatabase{}
	userDb.Start()
	defer userDb.Close()

	data, err := userDb.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
func (UserController) GetUserApps(c echo.Context) (err error) {

	userId := c.Param("userId")
	id, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userDb := database.AppDatabase{}
	userDb.Start()
	defer userDb.Close()

	data, err := userDb.GetByUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
