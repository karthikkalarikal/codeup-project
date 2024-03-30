package interfaces

import (
	"authentication/pkg/utils/request"
	"context"
)

type PaymentUsecase interface {
	GetPaymentIntent(ctx context.Context, req request.Payment) ([]byte, error)
}
