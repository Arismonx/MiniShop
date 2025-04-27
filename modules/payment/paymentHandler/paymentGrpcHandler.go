package paymentHandler

import _paymentUsecase "github.com/Arismonx/MiniShop/modules/payment/paymentUsecase"

//create interface and struct
type (
	PaymentGrpcHandlerService interface{}

	paymentGrpcHandler struct {
		paymentUsecase _paymentUsecase.PaymentUsecaseService
	}
)

//create constructor
func NewPaymentGrpcHandler(
	paymentUsecase _paymentUsecase.PaymentUsecaseService,
) PaymentGrpcHandlerService {
	return &paymentGrpcHandler{paymentUsecase}
}
