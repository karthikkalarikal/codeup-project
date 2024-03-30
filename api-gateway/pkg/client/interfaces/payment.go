package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
)

type PaymentClient interface {
	Payment(e echo.Context, in request.Stripe) ([]byte, error)
}
