package request

type UserSignUpRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

type UserSignInRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
