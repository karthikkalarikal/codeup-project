package request

type ForgotPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type NewPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}
