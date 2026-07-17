package main

import (
	"gamestore/controllers"
	"gamestore/database"
)

func main() {

	database.ConnectDatabase()

	defer database.DB.Close()

	controllers.StartMenu()

}
