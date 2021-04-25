package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/models"
)

func AddAlarm(c echo.Context) error {
	id_difabel := c.FormValue("id_difabel")
	id_perawat := c.FormValue("id_perawat")
	title := c.FormValue("title")
	jam := c.FormValue("jam")

	result, err := models.AddAlarm(id_difabel, id_perawat, title, jam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func UpdateAlarm(c echo.Context) error {
	id := c.FormValue("id")
	id_difabel := c.FormValue("id_difabel")
	id_perawat := c.FormValue("id_perawat")
	title := c.FormValue("title")
	jam := c.FormValue("jam")

	result, err := models.UpdateAlarm(id, id_difabel, id_perawat, title, jam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func DeleteAlarm(c echo.Context) error {
	id := c.Param("id")

	result, err := models.DeleteAlarm(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func FetchAlarm(c echo.Context) error {
	id := c.Param("id")

	result, err := models.FetchAlarm(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}
