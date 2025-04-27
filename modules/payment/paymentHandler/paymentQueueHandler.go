package paymentHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_paymentUsecase "github.com/Arismonx/MiniShop/modules/payment/paymentUsecase"
)

// create interface and struct
type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase _paymentUsecase.PaymentUsecaseService
	}
)

// create constructor
func NewPaymentQueueHandler(
	cfg *config.Config,
	paymentUsecase _paymentUsecase.PaymentUsecaseService,
) PaymentQueueHandlerService {
	return &paymentQueueHandler{cfg, paymentUsecase}
}
