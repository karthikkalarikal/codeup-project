package response

// type StripePayload struct {
// 	Currency string `json:"currency"`
// 	Amount   int    `json:"int"`
// }

type PaymentResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}
