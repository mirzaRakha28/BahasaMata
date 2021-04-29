package main

import (
	"os"

	"github.com/mirzaRakha28/BahasaMata/db"
	"github.com/mirzaRakha28/BahasaMata/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
