package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/helpers"
	"github.com/mirzaRakha28/BahasaMata/models"
)

func Register(c echo.Context) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	password = helpers.Md5Encrypt(password)
	typeData := c.FormValue("type")
	if typeData == "1" {

		result, err := models.RegisterDifabel(name, username, password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	} else {

		result, err := models.RegisterPerawat(name, username, password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}

}
