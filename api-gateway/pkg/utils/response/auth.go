package response

import (
	"time"
)

type UserSignUpResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Omit password in JSON responses
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"` // for proper representaion of null value in go sql.Null is used
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserSignInResponse struct {
	ID       int    `json:"id"`
	Admin    bool   `json:"admin"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Blocked  bool   `json:"blocked"`
}

type SignUpWrapper struct {
	User  UserSignUpResponse `json:"user"`
	Token string             `json:"token"`
}
