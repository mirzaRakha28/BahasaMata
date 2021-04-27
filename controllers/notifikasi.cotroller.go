package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/models"
)

func AddNotifikasi(c echo.Context) error {
	id_difabel := c.FormValue("id_difabel")
	id_perawat := c.FormValue("id_perawat")

	result, err := models.AddNotifikasi(id_difabel, id_perawat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}
