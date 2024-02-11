package domain

import "database/sql"

type User struct {
	ID        int            `json:"id"`
	Username  sql.NullString `json:"username"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"-"`    // Omit password in JSON responses
	Name      sql.NullString `json:"name"` // for proper representaion of null value in go sql.Null is used
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
