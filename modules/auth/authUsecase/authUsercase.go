package authUsecase

import (
	_authRepository "github.com/Arismonx/MiniShop/modules/auth/authRepository"
)

// *essential* Create interface and struct
type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository _authRepository.AuthRepositoryService
	}
)

// *essential* Create Constructor of struct authUsecase
func NewAuthUsecase(
	authRepository _authRepository.AuthRepositoryService,
) AuthUsecaseService {
	return &authUsecase{authRepository}
}
