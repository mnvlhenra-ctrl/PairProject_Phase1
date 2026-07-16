package controllers

import (
	"database/sql"
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

type ProductOrder struct {
	ProductID int
	Title     string
	Price     float64
	Stock     int
}

// ======================================
// BUY GAME
// ======================================

func BuyGame(session *models.Session) {

	utils.PrintTitle("BUY GAME")

	ShowProducts()

	productID := utils.InputInt("Product ID : ")

	qty := utils.InputInt("Quantity : ")

	query := `
	SELECT
		product_id,
		title,
		price,
		stock
	FROM products
	WHERE product_id = ?;
	`

	var product ProductOrder

	err := database.DB.QueryRow(
		query,
		productID,
	).Scan(
		&product.ProductID,
		&product.Title,
		&product.Price,
		&product.Stock,
	)

	if err == sql.ErrNoRows {

		fmt.Println()
		fmt.Println("Product not found.")
		return

	}

	if err != nil {

		fmt.Println(err)
		return

	}

	if qty <= 0 {

		fmt.Println("Quantity must be greater than zero.")
		return

	}

	if qty > product.Stock {

		fmt.Println()
		fmt.Println("Stock is not enough.")
		return

	}

	total := float64(qty) * product.Price

	utils.PrintLine()

	fmt.Println("Game  :", product.Title)
	fmt.Println("Price :", product.Price)
	fmt.Println("Qty   :", qty)
	fmt.Println("Total :", total)

	utils.PrintLine()

	confirm := utils.InputString("Confirm Purchase? (y/n) : ")

	if confirm != "y" && confirm != "Y" {

		fmt.Println()
		fmt.Println("Purchase cancelled.")
		return

	}

	// ======================================
	// START TRANSACTION
	// ======================================

	tx, err := database.DB.Begin()

	if err != nil {
		fmt.Println(err)
		return
	}

	// INSERT ORDER

	orderQuery := `
	INSERT INTO orders
	(user_id, status, total_price)
	VALUES
	(?, 'Paid', ?);
	`

	result, err := tx.Exec(
		orderQuery,
		session.UserID,
		total,
	)

	if err != nil {

		tx.Rollback()

		fmt.Println(err)

		return

	}

	orderID, err := result.LastInsertId()

	if err != nil {

		tx.Rollback()

		fmt.Println(err)

		return

	}

	// INSERT ORDER ITEM

	orderItemQuery := `
	INSERT INTO order_items
	(order_id, product_id, qty, price)
	VALUES
	(?,?,?,?);
	`

	_, err = tx.Exec(
		orderItemQuery,
		orderID,
		productID,
		qty,
		total,
	)

	if err != nil {

		tx.Rollback()

		fmt.Println(err)

		return

	}

	// UPDATE STOCK

	updateStock := `
	UPDATE products
	SET stock = stock - ?
	WHERE product_id = ?;
	`

	_, err = tx.Exec(
		updateStock,
		qty,
		productID,
	)

	if err != nil {

		tx.Rollback()

		fmt.Println(err)

		return

	}

	// COMMIT

	err = tx.Commit()

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("Purchase Successful!")
	fmt.Println("Order ID :", orderID)
	fmt.Println("Total    :", total)
	fmt.Println("===================================")

}

// ======================================
// MY ORDERS
// ======================================

func MyOrders(session *models.Session) {

	utils.PrintTitle("MY ORDERS")

	query := `
	SELECT
		o.order_id,
		p.title,
		oi.qty,
		oi.price,
		o.status,
		o.order_date
	FROM orders o
	JOIN order_items oi
	ON o.order_id = oi.order_id
	JOIN products p
	ON oi.product_id = p.product_id
	WHERE o.user_id = ?
	ORDER BY o.order_date DESC;
	`

	rows, err := database.DB.Query(
		query,
		session.UserID,
	)

	if err != nil {

		fmt.Println(err)
		return

	}

	defer rows.Close()

	fmt.Printf(
		"%-5s %-30s %-8s %-15s %-12s %-20s\n",
		"ID",
		"GAME",
		"QTY",
		"TOTAL",
		"STATUS",
		"ORDER DATE",
	)

	utils.PrintLine()

	found := false

	for rows.Next() {

		found = true

		var orderID int
		var title string
		var qty int
		var total float64
		var status string
		var orderDate string

		err := rows.Scan(
			&orderID,
			&title,
			&qty,
			&total,
			&status,
			&orderDate,
		)

		if err != nil {

			fmt.Println(err)
			return

		}

		fmt.Printf(
			"%-5d %-30s %-8d Rp %-12.0f %-12s %-20s\n",
			orderID,
			title,
			qty,
			total,
			status,
			orderDate,
		)

	}

	if !found {

		fmt.Println("No orders found.")

	}

	utils.PrintLine()

}
