package models

import "time"

type User struct {
	UserID    int
	Email     string
	Password  string
	Name      string
	Role      string
	CreatedAt time.Time
}
