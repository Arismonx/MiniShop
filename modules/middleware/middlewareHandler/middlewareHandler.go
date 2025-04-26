package middlewareHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_middlewareUsecase "github.com/Arismonx/MiniShop/modules/middleware/middlewareUsecase"
)

// Create interface and struct
type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase _middlewareUsecase.MiddlewareUsecaseService
	}
)

// Create constructor
func NewmiddlewareHandler(
	cfg *config.Config,
	middlewareUsecase _middlewareUsecase.MiddlewareUsecaseService,
) MiddlewareHandlerService {
	return &middlewareHandler{cfg, middlewareUsecase}
}
