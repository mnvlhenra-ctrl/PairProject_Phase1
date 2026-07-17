package controllers

import (
	"fmt"

	"gamestore/database"
	"gamestore/utils"
)

// ======================================
// REPORT MENU
// ======================================

func ReportMenu() {

	for {

		utils.PrintTitle("REPORT MENU")

		fmt.Println("1. Total Revenue")
		fmt.Println("2. Most Popular Game")
		fmt.Println("3. Product Sales Report")
		fmt.Println("4. Low Stock Alert")
		fmt.Println("0. Back")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:
			TotalRevenue()

		case 2:
			MostPopularGame()

		case 3:
			ProductSalesReport()

		case 4:
			LowStockAlert()

		case 0:
			return

		default:
			fmt.Println("Invalid Menu!")

		}

	}

}

// ======================================
// TOTAL REVENUE
// ======================================

func TotalRevenue() {

	utils.PrintTitle("TOTAL REVENUE")

	query := `
	SELECT
		IFNULL(SUM(total_price),0)
	FROM orders
	WHERE status='Paid';
	`

	var revenue float64

	err := database.DB.QueryRow(query).Scan(&revenue)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("Total Revenue :", utils.FormatRupiah(revenue))

}

// ======================================
// MOST POPULAR GAME
// ======================================

func MostPopularGame() {

	utils.PrintTitle("MOST POPULAR GAME")

	query := `
	SELECT
		p.title,
		SUM(oi.qty) AS total_sold
	FROM order_items oi
	JOIN products p
	ON oi.product_id = p.product_id
	GROUP BY p.product_id
	ORDER BY total_sold DESC
	LIMIT 1;
	`

	var title string
	var sold int

	err := database.DB.QueryRow(query).Scan(
		&title,
		&sold,
	)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("Game :", title)

	fmt.Println("Sold :", sold)

}

// ======================================
// PRODUCT SALES REPORT
// ======================================

func ProductSalesReport() {

	utils.PrintTitle("PRODUCT SALES REPORT")

	query := `
	SELECT
		p.product_id,
		p.title,
		IFNULL(SUM(oi.qty),0) AS total_sold,
		IFNULL(SUM(oi.price),0) AS revenue
	FROM products p
	LEFT JOIN order_items oi
	ON p.product_id = oi.product_id
	GROUP BY
		p.product_id,
		p.title
	ORDER BY total_sold DESC;
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	fmt.Printf(
		"%-5s %-30s %-10s %-15s\n",
		"ID",
		"GAME",
		"SOLD",
		"REVENUE",
	)

	utils.PrintLine()

	for rows.Next() {

		var id int
		var title string
		var sold int
		var revenue float64

		err := rows.Scan(
			&id,
			&title,
			&sold,
			&revenue,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(
			"%-5d %-30s %-10d %-15s\n",
			id,
			title,
			sold,
			utils.FormatRupiah(revenue),
		)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

}

// ======================================
// LOW STOCK ALERT
// ======================================

func LowStockAlert() {

	utils.PrintTitle("LOW STOCK ALERT")

	query := `
	SELECT
		product_id,
		title,
		stock
	FROM products
	WHERE stock < 15
	ORDER BY stock ASC;
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	fmt.Printf(
		"%-5s %-30s %-10s\n",
		"ID",
		"GAME",
		"STOCK",
	)

	utils.PrintLine()

	found := false

	for rows.Next() {

		found = true

		var id int
		var title string
		var stock int

		err := rows.Scan(
			&id,
			&title,
			&stock,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(
			"%-5d %-30s %-10d\n",
			id,
			title,
			stock,
		)

	}

	if !found {

		fmt.Println("No products need restocking.")

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

}
