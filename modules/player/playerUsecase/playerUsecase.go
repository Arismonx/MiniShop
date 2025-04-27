package playerUsecase

import _playerRepository "github.com/Arismonx/MiniShop/modules/player/playerRepository"

//Create interface and struct
type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository _playerRepository.PlayerRepositoryService
	}
)

//Create constructor
func NewPlayerUsecase(
	playerRepository _playerRepository.PlayerRepositoryService,
) PlayerUsecaseService {
	return &playerUsecase{playerRepository}
}
