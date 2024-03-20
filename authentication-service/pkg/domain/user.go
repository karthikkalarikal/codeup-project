package domain

import (
	"time"
)

type User struct {
	ID            int       `json:"id" gorm:"primarykey;autoIncrement"`
	Username      string    `json:"username" gorm:"unique"`
	Email         string    `json:"email" gorm:"not null"`
	VerifiedEmail bool      `json:"is_email_verified" gorm:"not null;default:false"`
	Password      string    `json:"-" gorm:"not null"` // Omit password in JSON responses
	FirstName     string    `json:"first_name" gorm:"not null"`
	LastName      string    `json:"last_name" gorm:"not null"` // for proper representaion of null value in go sql.Null is used
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null"`
	Admin         bool      `json:"isadmin" gorm:"default:false"`
	Blocked       bool      `json:"blocked" gorm:"default:false"`
}

type VerifyEmails struct {
	UserID     int       `json:"user_id" gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	Username   string    `json:"username" gorm:"not null;unique"`
	Email      string    `json:"email" gorm:"not null"`
	SecretCode string    `json:"secret_code" gorm:"not null"`
	IsUsed     bool      `json:"is_used" gorm:"not null;default:false"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;default:now()"`
	ExpiredAt  time.Time `json:"expired_at" gorm:"not null;default:now() + interval '15 minutes'"`
}
