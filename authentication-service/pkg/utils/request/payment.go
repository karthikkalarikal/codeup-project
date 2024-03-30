package request

type Payment struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
