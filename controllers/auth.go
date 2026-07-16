package controllers

import (
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
