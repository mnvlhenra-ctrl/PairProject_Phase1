package main

import (
	"fmt"

	"gamestore/controllers"
	"gamestore/database"
	"gamestore/utils"
)

func main() {

	database.ConnectDatabase()
	defer database.DB.Close()

	for {

		session, err := controllers.Login()

		if err != nil {

			fmt.Println()
			fmt.Println("Invalid email or password.")
			fmt.Println("Please try again.")
			fmt.Println()

			continue

		}

		utils.PrintTitle("WELCOME")

		fmt.Println("Hello,", session.Name)
		fmt.Println("Role :", session.Role)

		if session.Role == "admin" {

			controllers.AdminMenu(session)

		} else {

			controllers.CustomerMenu(session)

		}

		break

	}

	fmt.Println()
	fmt.Println("Thank you for using Game Store!")

}
