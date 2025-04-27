package paymentHandler

import _paymentUsecase "github.com/Arismonx/MiniShop/modules/payment/paymentUsecase"

//create interface and struct
type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		paymentUsecase _paymentUsecase.PaymentUsecaseService
	}
)

//create constructor
func NewPaymentHttpHandler(
	paymentUsecase _paymentUsecase.PaymentUsecaseService,
) PaymentHttpHandlerService {
	return &paymentHttpHandler{paymentUsecase}
}
