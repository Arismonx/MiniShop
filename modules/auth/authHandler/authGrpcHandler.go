package authHandler

import (
	_authUsecase "github.com/Arismonx/MiniShop/modules/auth/authUsecase"
)

// *essential* Create struct

type (
	authGrpcHandler struct {
		authUsecase _authUsecase.AuthUsecaseService
	}
)

// *essential* Create constructor of struct authGrpcHandler

func NewAuthGrpcHandler(
	authUsecase _authUsecase.AuthUsecaseService,
) _authUsecase.AuthUsecaseService {
	return &authGrpcHandler{authUsecase}
}
