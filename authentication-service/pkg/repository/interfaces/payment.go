package interfaces

import (
	"context"

	"github.com/stripe/stripe-go"
)

type PaymentRepo interface {
	Charge(ctx context.Context, currency string, amount int) (*stripe.PaymentIntent, string, error)
}
