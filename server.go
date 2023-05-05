package main

import (
	"github.com/AnggaArdhinata/k-stylehub/src/db"
	"github.com/AnggaArdhinata/k-stylehub/src/routers"
	
)

func main() {
	db.Init()

	e := routers.Init()

	e.Logger.Fatal(e.Start(":8080"))
}