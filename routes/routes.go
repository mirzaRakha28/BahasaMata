package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})
	e.POST("/register", controllers.Register)
	e.POST("/notifikasi", controllers.AddNotifikasi)
	e.POST("/pasien", controllers.AddPasien)
	e.POST("/login", controllers.Login)
	e.POST("/alarm", controllers.AddAlarm)
	e.POST("/alarm/perawat/:id", controllers.FetchAlarm)

	e.PUT("/alarm", controllers.UpdateAlarm)

	e.DELETE("/pasien", controllers.DeletePasien)
	e.DELETE("/notifikasi", controllers.DeleteNotifikasi)
	e.DELETE("/alarm/:id", controllers.DeleteAlarm)

	return e

}
