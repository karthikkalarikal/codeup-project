package request

type Stripe struct {
	Amount   int    `json:"amount" gorm:"not null"`
	Currency string `json:"currency"`
}
