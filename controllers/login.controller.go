package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/helpers"
	"github.com/mirzaRakha28/BahasaMata/models"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	typeData := c.FormValue("type")
	password = helpers.Md5Encrypt(password)
	if typeData == "1" {
		result, err := models.LoginDifabel(password, username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	} else {
		result, err := models.LoginPerawat(password, username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}

}
