package main

import (
	"github.com/mirzaRakha28/BahasaMata/db"
	"github.com/mirzaRakha28/BahasaMata/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
