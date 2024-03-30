package customErrors

import "github.com/stripe/stripe-go"

const (
	DatabaseTimeOut   = "database timeout"
	UserAlreadyExists = "user already exists"
	TimeOut           = "database query timed out"
)

func CardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "your card is expired"
	default:
		msg = "you card was declined"
	}
	return msg
}
