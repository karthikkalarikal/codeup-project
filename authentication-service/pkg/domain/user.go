package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primarykey;autoIncrement"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"` // Omit password in JSON responses
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"` // for proper representaion of null value in go sql.Null is used
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	Admin     bool      `json:"isadmin" gorm:"default:false"`
	Blocked   bool      `json:"blocked" gorm:"default:false"`
}
