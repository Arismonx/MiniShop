package server

import (
	_paymentHandler "github.com/Arismonx/MiniShop/modules/payment/paymentHandler"
	_paymentRepository "github.com/Arismonx/MiniShop/modules/payment/paymentRepository"
	_paymentUsecase "github.com/Arismonx/MiniShop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	//initialize payment
	repo := _paymentRepository.NewPaymentRepository(s.db)
	usecase := _paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := _paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := _paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	//Group path
	payment := s.app.Group("/payment_v1")

	//health check
	payment.GET("", s.healthCheck)
}
