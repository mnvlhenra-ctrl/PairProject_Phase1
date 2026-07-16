package models

type OrderItem struct {
	OrderItemID int
	OrderID     int
	ProductID   int
	Qty         int
	Price       float64
}