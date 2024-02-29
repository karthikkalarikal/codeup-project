package request

type UserSignUpRequest struct {
	Username        string `json:"username" validate:"required,min=6,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,min=8"`
	FirstName       string `json:"first_name" validate:"required,min=3"`
	LastName        string `json:"last_name" validate:"required,min=3"`
}
