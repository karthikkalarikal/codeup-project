package response

type UserSignInResponse struct {
	ID int `json:"id" copier:"must,nopanic"`

	Email string `json:"email" copier:"must,nopanic"`
}
