package controllers

import (
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

// ======================================
// LOGIN
// ======================================

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
	phone := utils.InputString("Phone    : ")

	// Mulai transaction.
	tx, err := database.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1) Buat user baru pada tabel users.
	result, err := tx.Exec(
		"INSERT INTO users (email, password, name, role) VALUES (?, ?, ?, 'customer');",
		email,
		password,
		name,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println()
		fmt.Println("Register failed.")
		fmt.Println("Email may already be used.")
		return
	}

	// Ambil user_id yang baru dibuat untuk relasi ke user_profiles.
	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// 2) Buat profil untuk user tersebut pada tabel user_profiles.
	//    user_name diisi sama dengan name, phone dari input pengguna.
	_, err = tx.Exec(
		"INSERT INTO user_profiles (user_id, user_name, phone) VALUES (?, ?, ?);",
		userID,
		name,
		phone,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// Semua langkah sukses -> simpan permanen.
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Register successful! Please login.")

}