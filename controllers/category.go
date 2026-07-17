package controllers

import (
	"fmt"

	"gamestore/database"
	"gamestore/models"
	"gamestore/utils"
)

// ======================================
// CATEGORY MENU
// ======================================

func CategoryMenu() {

	for {

		utils.PrintTitle("CATEGORY MENU")

		fmt.Println("1. Show Categories")
		fmt.Println("2. Add Category")
		fmt.Println("3. Update Category")
		fmt.Println("4. Delete Category")
		fmt.Println("0. Back")

		choice := utils.InputInt("Choose : ")

		switch choice {

		case 1:
			ShowCategories()

		case 2:
			AddCategory()

		case 3:
			UpdateCategory()

		case 4:
			DeleteCategory()

		case 0:
			return

		default:
			fmt.Println("Invalid Menu!")

		}

	}

}

// ======================================
// SHOW CATEGORIES
// ======================================

func ShowCategories() {

	query := `
	SELECT
		category_id,
		name,
		description
	FROM categories
	ORDER BY category_id;
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	utils.PrintTitle("CATEGORY LIST")

	fmt.Printf(
		"%-5s %-20s %-40s\n",
		"ID",
		"NAME",
		"DESCRIPTION",
	)

	utils.PrintLine()

	for rows.Next() {

		var category models.Category

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.Description,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(
			"%-5d %-20s %-40s\n",
			category.CategoryID,
			category.Name,
			category.Description,
		)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	utils.PrintLine()

}

// ======================================
// ADD CATEGORY
// ======================================

func AddCategory() {

	utils.PrintTitle("ADD CATEGORY")

	name := utils.InputString("Category Name : ")

	description := utils.InputString("Description : ")

	if name == "" {

		fmt.Println("Category name cannot be empty.")
		return

	}

	query := `
	INSERT INTO categories
	(name, description)
	VALUES
	(?, ?);
	`

	_, err := database.DB.Exec(
		query,
		name,
		description,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Category added successfully!")

}

// ======================================
// UPDATE CATEGORY
// ======================================

func UpdateCategory() {

	utils.PrintTitle("UPDATE CATEGORY")

	ShowCategories()

	id := utils.InputInt("Category ID : ")

	query := `
	SELECT
		category_id,
		name,
		description
	FROM categories
	WHERE category_id = ?;
	`

	var category models.Category

	err := database.DB.QueryRow(query, id).Scan(
		&category.CategoryID,
		&category.Name,
		&category.Description,
	)

	if err != nil {

		fmt.Println()
		fmt.Println("Category not found!")
		return

	}

	fmt.Println()

	utils.PrintLine()
	fmt.Println("Current Category")
	utils.PrintLine()

	fmt.Println("Name        :", category.Name)
	fmt.Println("Description :", category.Description)

	utils.PrintLine()

	name := utils.InputString("New Name : ")

	description := utils.InputString("New Description : ")

	if name == "" {

		fmt.Println("Category name cannot be empty.")
		return

	}

	updateQuery := `
	UPDATE categories
	SET
		name = ?,
		description = ?
	WHERE category_id = ?;
	`

	_, err = database.DB.Exec(
		updateQuery,
		name,
		description,
		id,
	)

	if err != nil {

		fmt.Println(err)
		return

	}

	fmt.Println()
	fmt.Println("Category updated successfully!")

}

// ======================================
// DELETE CATEGORY
// ======================================

func DeleteCategory() {

	utils.PrintTitle("DELETE CATEGORY")

	ShowCategories()

	id := utils.InputInt("Category ID : ")

	confirm := utils.InputString("Delete this category? (y/n) : ")

	if confirm != "y" && confirm != "Y" {

		fmt.Println()
		fmt.Println("Delete cancelled.")
		return

	}

	query := `
	DELETE FROM categories
	WHERE category_id = ?;
	`

	result, err := database.DB.Exec(
		query,
		id,
	)

	if err != nil {

		fmt.Println()
		fmt.Println("Cannot delete category.")
		fmt.Println("This category is still used by one or more products.")
		return

	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return
	}

	if rowsAffected == 0 {

		fmt.Println()
		fmt.Println("Category not found.")
		return

	}

	fmt.Println()
	fmt.Println("Category deleted successfully!")
}
