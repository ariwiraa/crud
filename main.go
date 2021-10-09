package main

import (
	"crud-echo/routes"
	"crud-echo/script/migration"
)

func main() {
	migration.Migrate()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
