package paymentUsecase

import _paymentRepository "github.com/Arismonx/MiniShop/modules/payment/paymentRepository"

//Create interface and struct
type (
	PaymentUsecaseService interface{}

	paymentUsecase struct {
		paymentRepository _paymentRepository.PaymentRepositoryService
	}
)

//Create constructor
func NewPaymentUsecase(
	paymentRepository _paymentRepository.PaymentRepositoryService,
) PaymentUsecaseService {
	return &paymentUsecase{paymentRepository}
}
