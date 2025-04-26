package authHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_authUsecase "github.com/Arismonx/MiniShop/modules/auth/authUsecase"
)

// *essential* Create interface and struct

type (
	AuthGrpcHandlerService interface{}

	authGrpcHandler struct {
		cfg         *config.Config
		authUsecase _authUsecase.AuthUsecaseService
	}
)

// *essential* Create constructor of struct authGrpcHandler

func NewAuthGrpcHandler(
	cfg *config.Config,
	authUsecase _authUsecase.AuthUsecaseService,
) AuthGrpcHandlerService {
	return &authGrpcHandler{cfg, authUsecase}
}
