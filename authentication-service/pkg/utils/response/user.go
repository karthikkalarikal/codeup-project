package response

type UserSignInResponse struct {
	ID       int    `json:"id" copier:"must,nopanic"`
	Password string `json:"password" copier:"must,nopanic"`
	Admin    bool   `json:"admin" copier:"must,nopanic"`
	Email    string `json:"email" copier:"must,nopanic"`
	Prime    bool   `json:"prime" copier:"must,nopanic"`
}
