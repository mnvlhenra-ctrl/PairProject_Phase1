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
// ======================================
// REGISTER
// ======================================

func Register() {

	utils.PrintTitle("REGISTER")

	name := utils.InputString("Name     : ")
	email := utils.InputString("Email    : ")
	password := utils.InputString("Password : ")

	query := `
	INSERT INTO users
	(email, password, name, role)
	VALUES
	(?, ?, ?, 'customer');
	`

	_, err := database.DB.Exec(
		query,
		email,
		password,
		name,
	)

	if err != nil {

		fmt.Println()
		fmt.Println("Register failed.")
		fmt.Println("Email may already be used.")
		return

	}

	fmt.Println()
	fmt.Println("Register successful! Please login.")

}