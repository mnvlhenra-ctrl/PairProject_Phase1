package models

import "time"

type Review struct {
	ReviewID  int
	UserID    int
	ProductID int
	Rating    int
	Comment   string
	CreatedAt time.Time
}
