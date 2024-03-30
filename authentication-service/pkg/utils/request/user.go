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

type AllUsers struct{}

type Search struct {
	SearchBy string `json:"seach_by"`
	Keyword  string `json:"keyword"`
}

type ForgotPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

