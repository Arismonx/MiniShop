package paymentHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_paymentUsecase "github.com/Arismonx/MiniShop/modules/payment/paymentUsecase"
)

// create interface and struct
type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase _paymentUsecase.PaymentUsecaseService
	}
)

// create constructor
func NewPaymentHttpHandler(
	cfg *config.Config,
	paymentUsecase _paymentUsecase.PaymentUsecaseService,
) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}
