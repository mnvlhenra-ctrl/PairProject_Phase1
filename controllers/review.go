package controllers

import (
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

// ======================================
// ADD REVIEW
// ======================================

func AddReview(session *models.Session) {

	// Tampilkan produk agar user tahu ID yang bisa diulas.
	ShowProducts()

	productID := utils.InputInt("Product ID : ")

	// Pastikan produk yang diulas memang ada.
	var title string

	err := database.DB.QueryRow(
		"SELECT title FROM products WHERE product_id = ?;",
		productID,
	).Scan(&title)

	if err != nil {
		fmt.Println()
		fmt.Println("Product not found!")
		return
	}

	rating := utils.InputInt("Rating (1-5) : ")

	// Validasi rentang rating.
	if rating < 1 || rating > 5 {
		fmt.Println()
		fmt.Println("Rating must be between 1 and 5.")
		return
	}

	comment := utils.InputString("Comment : ")

	// Simpan ulasan ke tabel reviews.
	_, err = database.DB.Exec(
		"INSERT INTO reviews (user_id, product_id, rating, comment) VALUES (?, ?, ?, ?);",
		session.UserID,
		productID,
		rating,
		comment,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Review added successfully!")
	fmt.Printf("Game    : %s\n", title)
	fmt.Printf("Rating  : %d/5\n", rating)

}