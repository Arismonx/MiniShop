package server

import (
	_authHandler "github.com/Arismonx/MiniShop/modules/auth/authHandler"
	_authRepository "github.com/Arismonx/MiniShop/modules/auth/authRepository"
	_authUsecase "github.com/Arismonx/MiniShop/modules/auth/authUsecase"
)

func (s *server) authService() {
	//initialize authService
	repo := _authRepository.NewAuthRepository(s.db)
	usecase := _authUsecase.NewAuthUsecase(repo)
	httpHandler := _authHandler.NewAuthHttpHandler(s.cfg, usecase)
	gRPCHandler := _authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = gRPCHandler

	//Group path
	auth := s.app.Group("/auth_v1")

	//health check
	auth.GET("", s.healthCheck)
}
