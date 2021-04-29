package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/BahasaMata/db"
	"github.com/mirzaRakha28/BahasaMata/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	a := echo.New()
	a.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})
	a.Logger.Fatal(e.Start(":3000"))
}
