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
	e.POST("/login", controllers.Login)
	return e

}
