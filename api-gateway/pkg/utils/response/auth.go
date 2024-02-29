package response

import "database/sql"

type UserSignUpResponse struct {
	ID        int            `json:"id"`
	Username  sql.NullString `json:"username"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"-"` // Omit password in JSON responses
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"` // for proper representaion of null value in go sql.Null is used
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
