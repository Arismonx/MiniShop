package authHandler

import (
	_authUsecase "github.com/Arismonx/MiniShop/modules/auth/authUsecase"
)

type (
	AuthGrpcHandlerService interface{}

	authGrpcHandler struct {
		authUsecase _authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(
	authUsecase _authUsecase.AuthUsecaseService,
) AuthGrpcHandlerService {
	return &authGrpcHandler{authUsecase}
}
