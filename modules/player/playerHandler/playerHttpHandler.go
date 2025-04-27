package playerHandler

import _playerUsecase "github.com/Arismonx/MiniShop/modules/player/playerUsecase"

//create interface and struct
type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		playerUsecase _playerUsecase.PlayerUsecaseService
	}
)

//create constructor
func NewPlayerHttpHandler(
	playerUsecase _playerUsecase.PlayerUsecaseService,
) PlayerHttpHandlerService {
	return &playerHttpHandler{playerUsecase}
}
