package usecase

import (
	"authentication/pkg/repository/interfaces"
	usecase "authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
	"encoding/json"
)

type paymentUsecaseImpl struct {
	// usecase interfaces.PaymentUsecase
	// cfg  config.Config
	repo interfaces.PaymentRepo
}

func NewPaymentUsecase(repo interfaces.PaymentRepo) usecase.PaymentUsecase {
	return &paymentUsecaseImpl{
		// cfg:  cfg,
		repo: repo,
	}
}

func (u *paymentUsecaseImpl) GetPaymentIntent(ctx context.Context, req request.Payment) ([]byte, error) {

	ok := true
	pi, msg, err := u.repo.Charge(ctx, req.Currency, req.Amount)
	if err != nil {
		ok = false

	}
	if ok {
		out, err := json.MarshalIndent(pi, "", " ")
		if err != nil {
			return nil, err
		}

		return out, nil
	} else {
		response := response.PaymentResponse{
			OK:      ok,
			Message: msg,
			Content: "",
		}
		out, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return nil, err
		}

		return out, nil
	}
}
