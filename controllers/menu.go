package controllers

import (
	"fmt"

	"gamestore/models"
	"gamestore/utils"
)

// ADMIN MENU

func AdminMenu(session *models.Session) {

	for {

		utils.PrintTitle("ADMIN MENU")

		fmt.Println("Welcome,", session.Name)
		fmt.Println("Role :", session.Role)
		utils.PrintLine()

		fmt.Println("1. Manage Products")
		fmt.Println("2. Manage Categories")
		fmt.Println("3. Reports")
		fmt.Println("4. Manage Reviews")
		fmt.Println("0. Logout")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:
			ProductMenu()

		case 2:
			CategoryMenu()

		case 3:
			ReportMenu()

		case 4:
			fmt.Println(">> Review Menu")

		case 0:
			fmt.Println()
			fmt.Println("Logout...")
			return

		default:
			fmt.Println("Invalid Menu!")

		}

	}

}

// CUSTOMER MENU

func CustomerMenu(session *models.Session) {

	for {

		utils.PrintTitle("CUSTOMER MENU")

		fmt.Println("Welcome,", session.Name)
		fmt.Println("Role :", session.Role)
		utils.PrintLine()

		fmt.Println("1. Browse Games")
		fmt.Println("2. Buy Game")
		fmt.Println("3. My Orders")
		fmt.Println("4. Add Review")
		fmt.Println("0. Logout")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:
			ShowProducts()

		case 2:
			BuyGame(session)

		case 3:
			MyOrders(session)

		case 4:
			AddReview(session)

		case 0:
			fmt.Println()
			fmt.Println("Logout...")
			return

		default:
			fmt.Println("Invalid Menu!")

		}

	}

}
