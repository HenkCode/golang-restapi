package main

import (
	"github.com/HenkCode/golang-restapi/db"
	"github.com/HenkCode/golang-restapi/routes"
)

func main() {
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}