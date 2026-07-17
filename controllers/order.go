package controllers

import (
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

// ======================================
// BUY GAME
// ======================================

func BuyGame(session *models.Session) {

	// Tampilkan daftar produk agar user tahu ID yang tersedia.
	ShowProducts()

	productID := utils.InputInt("Product ID : ")
	qty := utils.InputInt("Quantity   : ")

	// Validasi jumlah pembelian.
	if qty <= 0 {
		fmt.Println()
		fmt.Println("Quantity must be at least 1.")
		return
	}

	// Ambil data produk yang dipilih (judul, harga, stok).
	var product models.Product

	err := database.DB.QueryRow(
		"SELECT product_id, title, price, stock FROM products WHERE product_id = ?;",
		productID,
	).Scan(
		&product.ProductID,
		&product.Title,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		fmt.Println()
		fmt.Println("Product not found!")
		return
	}

	// Pastikan stok mencukupi.
	if product.Stock < qty {
		fmt.Println()
		fmt.Printf("Insufficient stock. Available: %d\n", product.Stock)
		return
	}

	totalPrice := product.Price * float64(qty)

	// Mulai transaction.
	tx, err := database.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1) Buat header order pada tabel orders. Status langsung "paid".
	result, err := tx.Exec(
		"INSERT INTO orders (user_id, status, total_price) VALUES (?, 'paid', ?);",
		session.UserID,
		totalPrice,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// Ambil order_id yang baru dibuat untuk dipakai pada order_items.
	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// 2) Buat detail order pada tabel order_items.
	_, err = tx.Exec(
		"INSERT INTO order_items (order_id, product_id, qty, price) VALUES (?, ?, ?, ?);",
		orderID,
		productID,
		qty,
		product.Price,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// 3) Kurangi stok produk sebanyak qty yang dibeli.
	_, err = tx.Exec(
		"UPDATE products SET stock = stock - ? WHERE product_id = ?;",
		qty,
		productID,
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
	fmt.Println("Purchase successful!")
	fmt.Printf("Order ID    : %d\n", orderID)
	fmt.Printf("Game        : %s\n", product.Title)
	fmt.Printf("Quantity    : %d\n", qty)
	fmt.Printf("Total Price : Rp %.0f\n", totalPrice)

}

// ======================================
// MY ORDERS
// ======================================

func MyOrders(session *models.Session) {

	query := `
	SELECT
		o.order_id,
		o.order_date,
		o.status,
		p.title,
		oi.qty,
		oi.price,
		o.total_price
	FROM orders o
	JOIN order_items oi ON o.order_id = oi.order_id
	JOIN products p     ON oi.product_id = p.product_id
	WHERE o.user_id = ?
	ORDER BY o.order_id;
	`

	rows, err := database.DB.Query(query, session.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	utils.PrintTitle("MY ORDERS")

	fmt.Printf(
		"%-8s %-12s %-8s %-25s %-5s %-15s\n",
		"ORDER",
		"DATE",
		"STATUS",
		"GAME",
		"QTY",
		"SUBTOTAL",
	)

	utils.PrintLine()

	found := false

	for rows.Next() {

		var (
			orderID    int
			orderDate  string
			status     string
			title      string
			qty        int
			price      float64
			totalPrice float64
		)

		err := rows.Scan(
			&orderID,
			&orderDate,
			&status,
			&title,
			&qty,
			&price,
			&totalPrice,
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		found = true

		
		if len(orderDate) > 10 {
			orderDate = orderDate[:10]
		}

		fmt.Printf(
			"%-8d %-12s %-8s %-25s %-5d Rp %-12.0f\n",
			orderID,
			orderDate,
			status,
			title,
			qty,
			price*float64(qty),
		)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	if !found {
		fmt.Println("You have no orders yet.")
	}

	utils.PrintLine()

}