package models

import "time"

type Order struct {
	OrderID    int
	UserID     int
	OrderDate  time.Time
	Status     string
	TotalPrice float64
}
