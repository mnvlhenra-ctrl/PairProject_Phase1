package controllers

import (
	"fmt"
	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

func Login() (*models.Session, error) {

	utils.PrintTitle("LOGIN")

	email := utils.InputString("Email    : ")

	password := utils.InputString("Password : ")

	session := models.Session{}

	query := `
	SELECT
		user_id,
		name,
		role
	FROM users
	WHERE email = ?
	AND password = ?;
	`

	err := database.DB.QueryRow(
		query,
		email,
		password,
	).Scan(
		&session.UserID,
		&session.Name,
		&session.Role,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil

}

func Register() {

	utils.PrintTitle("REGISTER")

	name := utils.InputString("Full Name : ")
	username := utils.InputString("Username  : ")
	phone := utils.InputString("Phone     : ")
	email := utils.InputString("Email     : ")
	password := utils.InputString("Password : ")

	tx, err := database.DB.Begin()

	if err != nil {
		fmt.Println(err)
		return
	}

	userQuery := `
	INSERT INTO users
	(email,password,name,role)
	VALUES
	(?,?,?,'customer');
	`

	result, err := tx.Exec(
		userQuery,
		email,
		password,
		name,
	)

	if err != nil {

		tx.Rollback()

		fmt.Println()
		fmt.Println("Register failed.")
		fmt.Println("Email may already exist.")
		return

	}

	userID, err := result.LastInsertId()

	if err != nil {

		tx.Rollback()
		fmt.Println(err)
		return

	}

	profileQuery := `
	INSERT INTO user_profiles
	(user_id,user_name,phone)
	VALUES
	(?,?,?);
	`

	_, err = tx.Exec(
		profileQuery,
		userID,
		username,
		phone,
	)

	if err != nil {

		tx.Rollback()

		fmt.Println(err)
		return

	}

	err = tx.Commit()

	if err != nil {

		fmt.Println(err)
		return

	}

	fmt.Println()
	fmt.Println("Register successful!")
	fmt.Println("Please login.")

}

// ======================================
// START MENU
// ======================================

func StartMenu() {

	for {

		utils.PrintTitle("GAME STORE")

		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:

			session, err := Login()

			if err != nil {

				fmt.Println()
				fmt.Println("Invalid email or password.")
				utils.Pause()
				continue

			}

			fmt.Println()
			fmt.Println("Welcome,", session.Name)

			if session.Role == "admin" {

				AdminMenu(session)

			} else {

				CustomerMenu(session)

			}

		case 2:

			Register()

		case 0:

			fmt.Println()
			fmt.Println("Thank you for using Game Store!")
			return

		default:

			fmt.Println()
			fmt.Println("Invalid Menu!")

		}

	}

}
