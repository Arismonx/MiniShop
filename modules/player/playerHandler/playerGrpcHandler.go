package playerHandler

import _playerUsecase "github.com/Arismonx/MiniShop/modules/player/playerUsecase"

//create interface and struct
type (
	PlayerGrpcHandlerService interface{}

	playerGrpcHandler struct {
		playerUsecase _playerUsecase.PlayerUsecaseService
	}
)

//create constructor
func NewPlayerGrpcHandler(
	playerUsecase _playerUsecase.PlayerUsecaseService,
) PlayerGrpcHandlerService {
	return &playerGrpcHandler{playerUsecase}
}
