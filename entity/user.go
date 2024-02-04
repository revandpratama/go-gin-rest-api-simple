package entity

import "time"

type User struct {
	Id        int
	Name      string
	Username  string
	Email     string
	Password  string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
