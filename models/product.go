package models

import "time"

type Product struct {
	ProductID  int
	CategoryID int
	Title      string
	Price      float64
	Stock      int
	CreatedAt  time.Time
}
