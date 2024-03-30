package repository

import (
	"authentication/pkg/config"
	"authentication/pkg/repository/interfaces"
	customErrors "authentication/pkg/utils/errors"
	"context"
	"fmt"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"gorm.io/gorm"
)

type paymentRepoImpl struct {
	cfg *config.StripeConfig
	DB  *gorm.DB
}

func NewPaymentRepo(DB *gorm.DB, cfg *config.StripeConfig) interfaces.PaymentRepo {
	return &paymentRepoImpl{
		cfg: cfg,
		DB:  DB,
	}
}

func (p *paymentRepoImpl) Charge(ctx context.Context, currency string, amount int) (*stripe.PaymentIntent, string, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.CreatePaymentIntent(currency, amount)
}

// payment intent
func (p *paymentRepoImpl) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = p.cfg.StripeSecret

	// customer_name := "John Doe"
	// address := &stripe.Address{
	// 	City:       "New York",
	// 	Country:    "US",
	// 	Line1:      "123 Main Street",
	// 	Line2:      "",
	// 	PostalCode: "123456",
	// 	State:      "New York",
	// }

	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(int64(amount)),
		Currency:    stripe.String(currency),
		Description: stripe.String("Export of software services"),
		// Customer:    stripe.String(),
		Shipping: &stripe.ShippingDetailsParams{
			Name: stripe.String("John Doe"),
			Address: &stripe.AddressParams{
				Line1:      stripe.String("510 Townsend St"),
				PostalCode: stripe.String("98140"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
			},
		},
	}

	// params.AddMetadata("key", "value")

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if sripErr, ok := err.(*stripe.Error); ok {
			msg = customErrors.CardErrorMessage(sripErr.Code)
		}
		fmt.Println("err", err)
		return nil, msg, err
	}

	fmt.Println("error ", err, stripe.String(pi.ID))
	return pi, ``, nil
}
