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

	utils.PrintTitle("ADD REVIEW")

	ShowProducts()

	productID := utils.InputInt("Product ID : ")

	rating := utils.InputInt("Rating (1-5) : ")

	comment := utils.InputString("Comment : ")

	// ==========================
	// VALIDATE RATING
	// ==========================

	if rating < 1 || rating > 5 {

		fmt.Println()
		fmt.Println("Rating must be between 1 and 5.")
		return

	}

	// ==========================
	// CHECK PURCHASE HISTORY
	// ==========================

	var purchased int

	purchaseQuery := `
	SELECT COUNT(*)
	FROM orders o
	JOIN order_items oi
	ON o.order_id = oi.order_id
	WHERE
		o.user_id = ?
	AND
		oi.product_id = ?;
	`

	err := database.DB.QueryRow(
		purchaseQuery,
		session.UserID,
		productID,
	).Scan(&purchased)

	if err != nil {

		fmt.Println(err)
		return

	}

	if purchased == 0 {

		fmt.Println()
		fmt.Println("You have never purchased this game.")
		return

	}

	// ==========================
	// CHECK DUPLICATE REVIEW
	// ==========================

	var reviewed int

	reviewQuery := `
	SELECT COUNT(*)
	FROM reviews
	WHERE
		user_id = ?
	AND
		product_id = ?;
	`

	err = database.DB.QueryRow(
		reviewQuery,
		session.UserID,
		productID,
	).Scan(&reviewed)

	if err != nil {

		fmt.Println(err)
		return

	}

	if reviewed > 0 {

		fmt.Println()
		fmt.Println("You have already reviewed this game.")
		return

	}

	// ==========================
	// INSERT REVIEW
	// ==========================

	insertQuery := `
	INSERT INTO reviews
	(
		user_id,
		product_id,
		rating,
		comment
	)
	VALUES
	(
		?,?,?,?
	);
	`

	_, err = database.DB.Exec(
		insertQuery,
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

}

// ======================================
// SHOW ALL REVIEWS
// ======================================

func ShowReviews() {

	utils.PrintTitle("GAME REVIEWS")

	query := `
	SELECT
		r.review_id,
		u.name,
		p.title,
		r.rating,
		r.comment
	FROM reviews r
	JOIN users u
	ON r.user_id = u.user_id
	JOIN products p
	ON r.product_id = p.product_id
	ORDER BY r.review_id;
	`

	rows, err := database.DB.Query(query)

	if err != nil {

		fmt.Println(err)
		return

	}

	defer rows.Close()

	fmt.Printf(
		"%-5s %-20s %-30s %-8s %-30s\n",
		"ID",
		"USER",
		"GAME",
		"RATE",
		"COMMENT",
	)

	utils.PrintLine()

	for rows.Next() {

		var reviewID int
		var user string
		var game string
		var rating int
		var comment string

		err := rows.Scan(
			&reviewID,
			&user,
			&game,
			&rating,
			&comment,
		)

		if err != nil {

			fmt.Println(err)
			return

		}

		fmt.Printf(
			"%-5d %-20s %-30s %-8d %-30s\n",
			reviewID,
			user,
			game,
			rating,
			comment,
		)

	}

	if err := rows.Err(); err != nil {

		fmt.Println(err)

	}

	utils.PrintLine()

}
