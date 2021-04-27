package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/models"
)

func AddPasien(c echo.Context) error {
	id_difabel := c.FormValue("id_difabel")
	id_perawat := c.FormValue("id_perawat")

	result, err := models.AddPasien(id_difabel, id_perawat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func DeletePasien(c echo.Context) error {
	id_difabel := c.FormValue("id_difabel")
	id_perawat := c.FormValue("id_perawat")

	result, err := models.DeletePasien(id_difabel, id_perawat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}
