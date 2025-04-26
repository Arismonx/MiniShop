package authHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_authUsecase "github.com/Arismonx/MiniShop/modules/auth/authUsecase"
)

// *essential* Create interface and struct
type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase _authUsecase.AuthUsecaseService
	}
)

// *essential* Create Constructor of struct authHttpHandler
func NewAuthHttpHandler(
	cfg *config.Config,
	authUsecase _authUsecase.AuthUsecaseService,
) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}
