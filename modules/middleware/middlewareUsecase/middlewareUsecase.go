package middlewareUsecase

import _middlewareRepository "github.com/Arismonx/MiniShop/modules/middleware/middlewareRepository"

// Create interface and struct
type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct {
		middlewareRepository _middlewareRepository.MiddlewareRepositoryService
	}
)

// Create constructor
func NewmiddlewareUsecase(
	middlewareRepository _middlewareRepository.MiddlewareRepositoryService,
) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}
