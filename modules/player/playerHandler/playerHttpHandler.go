package playerHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_playerUsecase "github.com/Arismonx/MiniShop/modules/player/playerUsecase"
)

// create interface and struct
type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase _playerUsecase.PlayerUsecaseService
	}
)

// create constructor
func NewPlayerHttpHandler(
	cfg *config.Config,
	playerUsecase _playerUsecase.PlayerUsecaseService,
) PlayerHttpHandlerService {
	return &playerHttpHandler{cfg, playerUsecase}
}
