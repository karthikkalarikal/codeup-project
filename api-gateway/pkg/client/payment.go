package client

import (
	"log"

	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	rpc "github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
)

type paymentClientImpl struct {
	rpc rpc.AuthService
}

func NewPaymentClient(rpc rpc.AuthService) interfaces.PaymentClient {
	return &paymentClientImpl{
		rpc: rpc,
	}
}

func (p *paymentClientImpl) Payment(e echo.Context, in request.Stripe) ([]byte, error) {
	body, err := p.rpc.Payment(e, in)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, err
}
