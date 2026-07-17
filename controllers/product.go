package controllers

import (
	"database/sql"
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

// ======================================
// PRODUCT MENU
// ======================================

func ProductMenu() {

	for {

		utils.PrintTitle("PRODUCT MENU")

		fmt.Println("1. Show Products")
		fmt.Println("2. Add Product")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("0. Back")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:
			ShowProducts()

		case 2:
			AddProduct()

		case 3:
			UpdateProduct()

		case 4:
			DeleteProduct()

		case 0:
			return

		default:
			fmt.Println("Invalid Menu!")

		}

	}

}

// ======================================
// SHOW PRODUCTS
// ======================================

func ShowProducts() {

	query := `
	SELECT
		p.product_id,
		p.title,
		c.name,
		p.price,
		p.stock
	FROM products p
	JOIN categories c
	ON p.category_id = c.category_id
	ORDER BY p.product_id;
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	utils.PrintTitle("PRODUCT LIST")

	fmt.Printf(
		"%-5s %-30s %-15s %-15s %-10s\n",
		"ID",
		"TITLE",
		"CATEGORY",
		"PRICE",
		"STOCK",
	)

	utils.PrintLine()

	for rows.Next() {

		var product models.Product
		var category string

		err := rows.Scan(
			&product.ProductID,
			&product.Title,
			&category,
			&product.Price,
			&product.Stock,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(
			"%-5d %-30s %-15s %-15s %-10d\n",
			product.ProductID,
			product.Title,
			category,
			utils.FormatRupiah(product.Price),
			product.Stock,
		)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	utils.PrintLine()

}

// ======================================
// ADD PRODUCT
// ======================================

func AddProduct() {

	utils.PrintTitle("ADD PRODUCT")

	ShowCategories()

	categoryID := utils.InputInt("Category ID : ")

	title := utils.InputString("Title : ")

	price := utils.InputFloat("Price : ")

	stock := utils.InputInt("Stock : ")

	if title == "" {

		fmt.Println("Title cannot be empty.")
		return

	}

	if price < 0 {

		fmt.Println("Price cannot be negative.")
		return

	}

	if stock < 0 {

		fmt.Println("Stock cannot be negative.")
		return

	}

	var temp int

	err := database.DB.QueryRow(
		"SELECT category_id FROM categories WHERE category_id = ?",
		categoryID,
	).Scan(&temp)

	if err == sql.ErrNoRows {

		fmt.Println("Category not found.")
		return

	}

	if err != nil {

		fmt.Println(err)
		return

	}

	query := `
	INSERT INTO products
	(category_id,title,price,stock)
	VALUES
	(?,?,?,?);
	`

	_, err = database.DB.Exec(
		query,
		categoryID,
		title,
		price,
		stock,
	)

	if err != nil {

		fmt.Println(err)
		return

	}

	fmt.Println()
	fmt.Println("Product added successfully!")

}

// ======================================
// UPDATE PRODUCT
// ======================================

func UpdateProduct() {

	utils.PrintTitle("UPDATE PRODUCT")

	ShowProducts()

	id := utils.InputInt("Product ID : ")

	query := `
	SELECT
		p.product_id,
		p.title,
		c.name,
		p.price,
		p.stock
	FROM products p
	JOIN categories c
	ON p.category_id = c.category_id
	WHERE p.product_id = ?;
	`

	var product models.Product
	var category string

	err := database.DB.QueryRow(query, id).Scan(
		&product.ProductID,
		&product.Title,
		&category,
		&product.Price,
		&product.Stock,
	)

	if err != nil {

		fmt.Println("Product not found!")
		return

	}

	utils.PrintLine()

	fmt.Println("Current Product")

	utils.PrintLine()

	fmt.Println("Title     :", product.Title)
	fmt.Println("Category  :", category)
	fmt.Println("Price     :", utils.FormatRupiah(product.Price))
	fmt.Println("Stock     :", product.Stock)

	utils.PrintLine()

	ShowCategories()

	categoryID := utils.InputInt("New Category ID : ")

	var temp int

	err = database.DB.QueryRow(
		"SELECT category_id FROM categories WHERE category_id = ?",
		categoryID,
	).Scan(&temp)

	if err == sql.ErrNoRows {

		fmt.Println("Category not found.")
		return

	}

	if err != nil {

		fmt.Println(err)
		return

	}

	title := utils.InputString("New Title : ")

	price := utils.InputFloat("New Price : ")

	stock := utils.InputInt("New Stock : ")

	if title == "" {

		fmt.Println("Title cannot be empty.")
		return

	}

	updateQuery := `
	UPDATE products
	SET
		category_id = ?,
		title = ?,
		price = ?,
		stock = ?
	WHERE product_id = ?;
	`

	_, err = database.DB.Exec(
		updateQuery,
		categoryID,
		title,
		price,
		stock,
		id,
	)

	if err != nil {

		fmt.Println(err)
		return

	}

	fmt.Println()
	fmt.Println("Product updated successfully!")

}

// ======================================
// DELETE PRODUCT
// ======================================

func DeleteProduct() {

	utils.PrintTitle("DELETE PRODUCT")

	ShowProducts()

	id := utils.InputInt("Product ID : ")

	confirm := utils.InputString("Delete this product? (y/n) : ")

	if confirm != "y" && confirm != "Y" {

		fmt.Println()
		fmt.Println("Delete cancelled.")
		return

	}

	query := `
	DELETE FROM products
	WHERE product_id = ?;
	`

	_, err := database.DB.Exec(
		query,
		id,
	)

	if err != nil {

		fmt.Println()
		fmt.Println("Cannot delete product.")
		fmt.Println("This product is already used in one or more orders.")
		return

	}

	fmt.Println()
	fmt.Println("Product deleted successfully!")

}
